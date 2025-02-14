package gapi


type DataDetails struct {
	Asin                 string  `json:"asin"`
	ProductTitle         string  `json:"product_title"`
	ProductPrice         string  `json:"product_price"`
	ProductOriginalPrice string  `json:"product_original_price"`
	Currency             string  `json:"currency"`
	Country              string  `json:"country"`
	ProductByline        string  `json:"product_byline"`
	ProductBylineLink    string  `json:"product_byline_link"`
	ProductStarRating    float64 `json:"product_star_rating"`
	ProductNumRatings    int     `json:"product_num_ratings"`
	ProductURL           string  `json:"product_url"`
	ProductPhoto         string  `json:"product_photo"`
	ProductNumOffers     int     `json:"product_num_offers"`
	ProductAvailability  string  `json:"product_availability"`
	IsBestSeller         bool    `json:"is_best_seller"`
	IsAmazonChoice       bool    `json:"is_amazon_choice"`
	IsPrime              bool    `json:"is_prime"`
	ClimatePledgeFriendly bool   `json:"climate_pledge_friendly"`
	SalesVolume          string  `json:"sales_volume"`
	ProductDescription   string  `json:"product_description"`
	VideoThumbnail       string  `json:"video_thumbnail"`
	HasVideo             bool    `json:"has_video"`
	CustomersSay         string  `json:"customers_say"`
	Delivery             string  `json:"delivery"`
	PrimaryDeliveryTime  string  `json:"primary_delivery_time"`
	HasAPlus             bool    `json:"has_aplus"`
	HasBrandStory        bool    `json:"has_brandstory"`
}