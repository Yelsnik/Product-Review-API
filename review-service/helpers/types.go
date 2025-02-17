package helpers

type DataDetails struct {
	Asin                  string `json:"asin"`
	ProductTitle          string `json:"product_title"`
	ProductPrice          string `json:"product_price"`
	ProductOriginalPrice  string `json:"product_original_price"`
	Currency              string `json:"currency"`
	Country               string `json:"country"`
	ProductByline         string `json:"product_byline"`
	ProductBylineLink     string `json:"product_byline_link"`
	ProductStarRating     string `json:"product_star_rating"`
	ProductNumRatings     int    `json:"product_num_ratings"`
	ProductURL            string `json:"product_url"`
	ProductPhoto          string `json:"product_photo"`
	ProductNumOffers      int    `json:"product_num_offers"`
	ProductAvailability   string `json:"product_availability"`
	IsBestSeller          bool   `json:"is_best_seller"`
	IsAmazonChoice        bool   `json:"is_amazon_choice"`
	IsPrime               bool   `json:"is_prime"`
	ClimatePledgeFriendly bool   `json:"climate_pledge_friendly"`
	SalesVolume           string `json:"sales_volume"`
	ProductDescription    string `json:"product_description"`
	VideoThumbnail        string `json:"video_thumbnail"`
	HasVideo              bool   `json:"has_video"`
	CustomersSay          string `json:"customers_say"`
	Delivery              string `json:"delivery"`
	PrimaryDeliveryTime   string `json:"primary_delivery_time"`
	HasAPlus              bool   `json:"has_aplus"`
	HasBrandStory         bool   `json:"has_brandstory"`
}

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

type ResponseDetails struct {
	Status     string      `json:"status"`
	RequestId  string      `json:"request_id"`
	Parameters any         `json:"parameters"`
	Data       DataDetails `json:"data"`
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
