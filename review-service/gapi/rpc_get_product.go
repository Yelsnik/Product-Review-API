package gapi

import (
	"context"
	"review-service/review"
	"review-service/val"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetProducts(ctx context.Context, req *review.GetProductsRequest) (*review.GetProductsResponse, error) {
	violations := validateGetProductsReq(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	result, err := server.helpers.GetAmazonProducts(req.GetPage())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get amazon products: %s", err)
	}

	var products []*review.Products
	for _, v := range result.Data.Products {
		results := &review.Products{
			Asin:                 v.Asin,
			ProductTitle:         v.ProductTitle,
			ProductPrice:         v.ProductPrice,
			ProductOriginalPrice: v.ProductOriginalPrice,
			Currency:             v.Currency,
			ProductUrl:           v.ProductUrl,
			ProductPhoto:         v.ProductPhoto,
		}
		products = append(products, results)
	}

	response := &review.GetProductsResponse{
		Product: products,
	}

	return response, nil
}

func validateGetProductsReq(req *review.GetProductsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetPage(), 2, 100); err != nil {
		violations = append(violations, fielViolation("page", err))
	}

	if err := val.ValidateString(req.GetCountry(), 2, 100); err != nil {
		violations = append(violations, fielViolation("country", err))
	}

	return violations
}
