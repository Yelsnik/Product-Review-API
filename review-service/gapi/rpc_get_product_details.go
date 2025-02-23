package gapi

import (
	"context"
	"fmt"
	"review-service/review"
	"review-service/val"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetProductDetails(ctx context.Context, req *review.GetProductDetailsRequest) (*review.GetProductDetailsResponse, error) {

	violations := validateGetProductDetailsRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	res, err := server.helpers.GetAmazonProductDetails(ctx, req.GetAsin(), req.GetCountry())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get amazon product: %s", err)
	}

	fmt.Println(res)

	result := *res

	response := &review.GetProductDetailsResponse{
		Product: &review.ProductDetails{
			Asin:                 result.Data.Asin,
			ProductTitle:         result.Data.ProductTitle,
			ProductPrice:         result.Data.ProductPrice,
			ProductOriginalPrice: result.Data.ProductOriginalPrice,
			Currency:             result.Data.Currency,
			Country:              result.Data.Country,
			ProductUrl:           result.Data.ProductURL,
			ProductPhoto:         result.Data.ProductPhoto,
			ProductAvailability:  result.Data.ProductAvailability,
		},
	}

	return response, nil
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
