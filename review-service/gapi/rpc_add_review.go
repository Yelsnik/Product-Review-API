package gapi

import (
	"context"
	"fmt"
	db "review-service/db/sqlc"
	"review-service/nlp"
	"review-service/review"
	"review-service/val"
	"sync"

	"github.com/jackc/pgx/v5"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) AddReview(ctx context.Context, req *review.AddReviewRequest) (*review.AddReviewResponse, error) {
	violations := validateAddReviewReq(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	errChan := make(chan error, 2)
	var sentiment *nlp.SentimentResponse

	// check the product id
	country := "US"

	go func() {
		defer wg.Done()
		_, err := server.helpers.GetAmazonProductDetails(req.GetProductId(), country)

		errChan <- err
	}()

	// analyze the sentiment
	go func() {
		defer wg.Done()
		s, err := server.client.Analyze(ctx, req.GetReview())

		errChan <- err
		sentiment = s
		fmt.Println(sentiment)
	}()

	fmt.Println(sentiment)

	// Wait for goroutines and check for errors
	wg.Wait()
	close(errChan)
	for err := range errChan {
		if err != nil {
			return nil, status.Errorf(codes.Internal, "%s", err)
		}
	}

	// add review to the db
	r, err := server.store.GetReviewByProductID(ctx, req.GetProductId())
	if err != nil {
		if err == pgx.ErrNoRows {
			// create a review
			arg := db.CreateReviewParams{
				ProductID:    req.GetProductId(),
				NumOfReviews: 0,
			}

			// create review
			r, err := server.store.CreateReview(ctx, arg)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to create review in the db: %s", err)
			}

			params := AddReviewResponseParams{
				server:    server,
				sentiment: sentiment,
				getreview: req.GetReview(),
				r:         r,
			}

			// add review message
			return AddReviewResponse(ctx, params)
		}
		return nil, status.Errorf(codes.Internal, "failed to get review: %s", err)
	}

	params := AddReviewResponseParams{
		server:    server,
		sentiment: sentiment,
		getreview: req.GetReview(),
		r:         r,
	}

	// add review message
	return AddReviewResponse(ctx, params)
}

type AddReviewResponseParams struct {
	server    *Server
	sentiment *nlp.SentimentResponse
	getreview string
	getproductid string
	r         db.Review
}

func AddReviewResponse(ctx context.Context, params AddReviewResponseParams) (*review.AddReviewResponse, error) {

	// add review message
	_, err := params.server.store.AddReviewTx(ctx, db.CreateReviewMessageParams{
		Review:   params.getreview,
		Score:    float64(params.sentiment.Score),
		Label:    params.sentiment.Label,
		ReviewID: params.r.ID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add review to database: %s", err)
	}

	// update leaderboard
	err = params.server.leaderboard.UpdateLeaderBoard(ctx, params.getproductid, float64(params.sentiment.Score))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update leaderboard: %s", err)
	}

	response := &review.AddReviewResponse{
		Message: "successfully added review",
	}

	return response, nil

}

func validateAddReviewReq(req *review.AddReviewRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetReview(), 2, 100); err != nil {
		violations = append(violations, fielViolation("review", err))
	}

	if err := val.ValidateString(req.GetProductId(), 2, 100); err != nil {
		violations = append(violations, fielViolation("product_id", err))
	}

	return violations
}
