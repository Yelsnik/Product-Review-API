package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"review-service/util"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Helpers interface {
	GetAmazonProducts(page string) (*Response, error)
	GetAmazonProductDetails(ctx context.Context, asin, country string) (*ResponseDetails, error)
}

type Helper struct {
	config util.Config
	redis  *redis.Client
}

func NewHelper(config util.Config, redis *redis.Client) Helpers {
	return &Helper{
		config: config,
		redis:  redis,
	}
}

func (h *Helper) GetAmazonProducts(page string) (*Response, error) {
	client := resty.New()

	url := fmt.Sprintf("%s?query=Phone&page=%s&country=US&sort_by=RELEVANCE&product_condition=ALL&is_prime=false&deals_and_discounts=NONE", h.config.RapidAPISearchUrl, page)

	resp, err := client.R().
		SetHeader("x-rapidapi-key", h.config.RapidAPIKey).
		SetHeader("x-rapidapi-host", h.config.RapidAPIHost).
		Get(url)

	var result Response
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmarshal response: %s", err)
	}

	return &result, err
}

func (h *Helper) GetAmazonProductDetails(ctx context.Context, asin, country string) (*ResponseDetails, error) {

	res, err := h.redis.HGet(ctx, asin, "details").Result()
	if err != nil {
		client := resty.New()

		url := fmt.Sprintf("%s?asin=%s&country=%s", h.config.RapidAPIDetailsUrl, asin, country)

		resp, err := client.R().
			SetHeader("x-rapidapi-key", h.config.RapidAPIKey).
			SetHeader("x-rapidapi-host", h.config.RapidAPIHost).
			Get(url)
		if err != nil {
			return nil, err
		}

		var result ResponseDetails
		err = json.Unmarshal(resp.Body(), &result)
		if err != nil {
			return nil, err
		}

		if result.Status == "ERROR" {
			return nil, fmt.Errorf("error getting the products")
		}

		// cache for next time
		productJSON, _ := json.Marshal(result)
		h.redis.HSet(ctx, result.Data.Asin, "details", productJSON)
		h.redis.Expire(ctx, result.Data.Asin, 24*time.Hour)

		return &result, err
	}

	var product ResponseDetails
	json.Unmarshal([]byte(res), &product)
	fmt.Println("Cache hit! Returning product details from Redis")
	return &product, nil
}
