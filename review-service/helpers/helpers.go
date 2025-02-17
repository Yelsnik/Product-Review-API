package helpers

import (
	"encoding/json"
	"fmt"
	"review-service/util"

	"github.com/go-resty/resty/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Helpers interface {
	GetAmazonProducts(page string) (*Response, error)
	GetAmazonProductDetails(asin, country string) (*ResponseDetails, error)
}

type Helper struct {
	config util.Config
}

func NewHelper(config util.Config) Helpers {
	return &Helper{
		config: config,
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

func (h *Helper) GetAmazonProductDetails(asin, country string) (*ResponseDetails, error) {
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

	return &result, err
}
