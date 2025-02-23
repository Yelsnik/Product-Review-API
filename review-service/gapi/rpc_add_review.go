package gapi

import (
	"context"
	"log"
	db "review-service/db/sqlc"
	"review-service/nlp"
	"review-service/review"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) AddReview(ctx context.Context, req *review.AddReviewRequest) (*review.AddReviewResponse, error) {
	start := time.Now()
	log.Println("Starting AddReview")

	violations := validateAddReviewReq(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	wg := &sync.WaitGroup{}
	wg.Add(3)

	errChan := make(chan error, 3)
	sentimentChan := make(chan *nlp.SentimentResponse, 1)
	reviewChan := make(chan db.Review, 1)

	// check the product id
	country := "US"

	startAmazon := time.Now()
	go func() {
		defer wg.Done()
		_, err := server.helpers.GetAmazonProductDetails(ctx, req.GetProductId(), country)
		if err != nil {
			errChan <- err
			return
		}
		log.Println("GetAmazonProductDetails took:", time.Since(startAmazon))
	}()

	// analyze the sentiment
	startSentiment := time.Now()
	go func() {
		defer wg.Done()
		s, err := server.client.Analyze(ctx, req.GetReview())
		if err != nil {
			errChan <- err
			return
		}
		sentimentChan <- s
		// fmt.Println(sentimentChan)
		log.Println("Analyze Sentiment took:", time.Since(startSentiment))
	}()

	// fetch review
	go func() {
		defer wg.Done()
		startReviewFetch := time.Now()
		r, err := server.store.GetReviewByProductID(ctx, req.GetProductId())
		log.Println("GetReviewByProductID took:", time.Since(startReviewFetch))
		if err != nil {
			errChan <- err
		}

		reviewChan <- r
	}()

	// Wait for goroutines and check for errors
	wg.Wait()
	close(errChan)
	close(sentimentChan)
	close(reviewChan)

	var sentiment *nlp.SentimentResponse
	var rev db.Review
	select {
	case sentiment = <-sentimentChan:
	default:
		return nil, status.Errorf(codes.Internal, "sentiment analysis failed")
	}
	rev = <-reviewChan

	for err := range errChan {
		if err != nil {
			if err == pgx.ErrNoRows {
				// create a review
				startReviewCreate := time.Now()
				arg := db.CreateReviewParams{
					ProductID:    req.GetProductId(),
					NumOfReviews: 0,
				}

				// create review
				r, err := server.store.CreateReview(ctx, arg)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "failed to create review in the db: %s", err)
				}
				log.Println("CreateReview took:", time.Since(startReviewCreate))

				params := AddReviewResponseParams{
					server:    server,
					sentiment: sentiment,
					getreview: req.GetReview(),
					r:         r,
				}

				// add review message
				return AddReviewResponse(ctx, params)

			}
			return nil, status.Errorf(codes.Internal, "%s", err)
		}
	}

	// add review to the db
	params := AddReviewResponseParams{
		server:    server,
		sentiment: sentiment,
		getreview: req.GetReview(),
		r:         rev,
	}

	log.Println("Total execution time:", time.Since(start))
	// add review message
	return AddReviewResponse(ctx, params)
}
