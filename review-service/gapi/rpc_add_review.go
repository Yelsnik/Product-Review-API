package gapi

import (
	"context"
	db "review-service/db/sqlc"
	"review-service/review"
	"review-service/val"

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

	// check the product id
	country := "US"

	_, err := server.helpers.GetAmazonProductDetails(req.GetProductId(), country)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id: %s", err)
	}

	// analyze the sentiment
	sentiment, err := server.client.Analyze(ctx, req.GetReview())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to analyze the review: %s", err)
	}

	// update leaderboard
	err = server.leaderboard.UpdateLeaderBoard(ctx, req.GetProductId(), float64(sentiment.Score))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update leaderboard: %s", err)
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

			// add review message
			_, err = server.store.AddReviewTx(ctx, db.CreateReviewMessageParams{
				Review:   req.GetReview(),
				Score:    float64(sentiment.Score),
				Label:    sentiment.Label,
				ReviewID: r.ID,
			})
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to add review message to the db: %s", err)
			}

			response := &review.AddReviewResponse{
				Message: "successfully added review",
			}

			return response, nil
		}
		return nil, status.Errorf(codes.Internal, "failed to get review: %s", err)
	}

	// add review message
	_, err = server.store.AddReviewTx(ctx, db.CreateReviewMessageParams{
		Review:   req.GetReview(),
		Score:    float64(sentiment.Score),
		Label:    sentiment.Label,
		ReviewID: r.ID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add review message to the db: %s", err)
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
