package helpers

import (
	"fmt"
	"review-service/util"

	"github.com/go-resty/resty/v2"
)

type Helpers interface {
	GetAmazonProducts(page string) (*resty.Response, error)
	GetAmazonProductDetails(asin, country string) (*resty.Response, error)
}

type Helper struct {
	config util.Config
}

func NewHelper(config util.Config) Helpers {
	return &Helper{
		config: config,
	}
}

func (h *Helper) GetAmazonProducts(page string) (*resty.Response, error) {
	client := resty.New()

	url := fmt.Sprintf("%s?query=Phone&page=%s&country=US&sort_by=RELEVANCE&product_condition=ALL&is_prime=false&deals_and_discounts=NONE", h.config.RapidAPISearchUrl, page)

	resp, err := client.R().
		SetHeader("x-rapidapi-key", h.config.RapidAPIKey).
		SetHeader("x-rapidapi-host", h.config.RapidAPIHost).
		Get(url)

	return resp, err
}

func (h *Helper) GetAmazonProductDetails(asin, country string) (*resty.Response, error) {
	client := resty.New()

	url := fmt.Sprintf("%s?asin=%s&country=%s", h.config.RapidAPIDetailsUrl, asin, country)

	resp, err := client.R().
		SetHeader("x-rapidapi-key", h.config.RapidAPIKey).
		SetHeader("x-rapidapi-host", h.config.RapidAPIHost).
		Get(url)

	return resp, err
}
