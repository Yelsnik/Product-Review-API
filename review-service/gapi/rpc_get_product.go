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

type AmazonProducts struct {
	Asin                     string `json:"asin"`
	ProductTitle             string `json:"product_title"`
	ProductPrice             string `json:"product_price"`
	ProductOriginalPrice     string `json:"product_original_price"`
	Currency                 string `json:"currency"`
	ProductStarRating        string `json:"product_star_rating"`
	ProductNumRatings        int64  `json:"product_num_ratings"`
	ProductUrl               string `json:"product_url"`
	ProductPhoto             string `json:"product_photo"`
	ProductNumOffers         int64  `json:"product_num_offers"`
	ProductMinimumOfferPrice string `json:"product_minimum_offer_price"`
	IsBestSeller             bool   `json:"is_best_seller"`
	IsAmazonChoice           bool   `json:"is_amazon_choice"`
	IsPrime                  bool   `json:"is_prime"`
	ClimatePledgeFriendly    bool   `json:"climate_pledge_friendly"`
	SalesVolume              string `json:"sales_volume"`
	Delivery                 string `json:"delivery"`
	HasVariations            bool   `json:"has_variations"`
}

type Data struct {
	TotalProduct int64            `json:"total_product"`
	Country      string           `json:"country"`
	Domain       string           `json:"domain"`
	Products     []AmazonProducts `json:"products"`
}

type Response struct {
	Status     string `json:"status"`
	Request_id string `json:"request_id"`
	Parameters any    `json:"parameters"`
	Data       Data   `json:"data"`
}

func (server *Server) GetProducts(ctx context.Context, req *review.GetProductsRequest) (*review.GetProductsResponse, error) {
	violations := validateGetProductsReq(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	resp, err := server.helpers.GetAmazonProducts(req.GetPage())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get amazon products: %s", err)
	}

	//fmt.Println(string(resp.Body()))

	var result Response
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmarshal response: %s", err)
	}

	// fmt.Println(result.Data.Products)

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

	// fmt.Println(products)

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
