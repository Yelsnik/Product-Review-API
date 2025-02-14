package gapi

import (
	"context"
	"encoding/json"
	"review-service/review"
	"review-service/val"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ResponseDetails struct {
	Status     string      `json:"status"`
	RequestId  string      `json:"request_id"`
	Parameters any         `json:"parameters"`
	Data       DataDetails `json:"data"`
}

func (server *Server) GetProductDetails(ctx context.Context, req *review.GetProductDetailsRequest) (*review.GetProductDetailsResponse, error) {

	resp, err := server.helpers.GetAmazonProductDetails(req.GetAsin(), req.GetCountry())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get amazon product: %s", err)
	}

	var result Response
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmarshal response: %s", err)
	}

	response := &review.GetProductDetailsResponse{
		Product: &review.ProductDetails{},
	}

	return nil, nil
}

func validateGetProductDetailsRequest(req *review.GetProductDetailsRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateString(req.GetAsin(), 2, 100); err != nil {
		violations = append(violations, fielViolation("asin", err))
	}

	if err := val.ValidateString(req.GetCountry(), 2, 100); err != nil {
		violations = append(violations, fielViolation("country", err))
	}

	return violations
}
