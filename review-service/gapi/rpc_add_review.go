package gapi

import (
	"context"
	"review-service/review"
	"review-service/val"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) AddReview(ctx context.Context, req *review.AddReviewRequest) (*review.AddReviewResponse, error) {
	violations := validateAddReviewReq(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	return nil, nil
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
