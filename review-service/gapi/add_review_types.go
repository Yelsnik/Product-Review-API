package gapi

import (
	"context"
	"log"
	db "review-service/db/sqlc"
	"review-service/nlp"
	"review-service/review"
	"review-service/val"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AddReviewResponseParams struct {
	server       *Server
	sentiment    *nlp.SentimentResponse
	getreview    string
	getproductid string
	r            db.Review
}

func AddReviewResponse(ctx context.Context, params AddReviewResponseParams) (*review.AddReviewResponse, error) {
	start := time.Now()

	startaddreviewtx := time.Now()
	// add review message
	_, err := params.server.store.AddReviewTx(ctx, db.CreateReviewMessageParams{
		Review:   params.getreview,
		Score:    float64(params.sentiment.Score),
		Label:    params.sentiment.Label,
		ReviewID: params.r.ID,
	})
	log.Println("addreviewtx took:", time.Since(startaddreviewtx))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add review to database: %s", err)
	}

	// update leaderboard
	startupdateleaderboard := time.Now()
	err = params.server.leaderboard.UpdateLeaderBoard(ctx, params.getproductid, float64(params.sentiment.Score))
	log.Println("updateleaderboard took:", time.Since(startupdateleaderboard))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update leaderboard: %s", err)
	}

	response := &review.AddReviewResponse{
		Message: "successfully added review",
	}

	log.Println("Total execution time for addreviewtxfunc:", time.Since(start))
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
