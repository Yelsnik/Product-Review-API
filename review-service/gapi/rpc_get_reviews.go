package gapi

import (
	"context"
	"fmt"
	"review-service/review"
	"review-service/val"

	"github.com/jackc/pgx/v5"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetReviews(ctx context.Context, req *review.GetReviewsRequest) (*review.GetReviewsResponse, error) {

	fmt.Println("1", req)
	violations := validateGetReviewsReq(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	r, err := server.store.GetReviewByProductID(ctx, req.GetProductId())
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "product has no review: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get reviews: %s", err)
	}

	reviews, err := server.store.GetReviewMessagesByReview(ctx, r.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get review messages: %s", err)
	}

	var rev []*review.Reviews

	for _, v := range reviews {
		data := &review.Reviews{
			Id:        v.ID.String(),
			Review:    v.Review,
			Score:     float32(v.Score),
			Label:     v.Label,
			ReviewId:  v.ReviewID.String(),
			CreatedAt: timestamppb.New(v.CreatedAt),
		}

		rev = append(rev, data)
	}

	response := &review.GetReviewsResponse{
		Reviews: rev,
	}

	return response, nil
}

func validateGetReviewsReq(req *review.GetReviewsRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateString(req.GetProductId(), 2, 100); err != nil {
		violations = append(violations, fielViolation("product_id", err))
	}

	return violations
}
