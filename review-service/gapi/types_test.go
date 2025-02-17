package gapi

import (
	"review-service/helpers"
	"review-service/leaderboard"
	"review-service/review"
	"review-service/util"
)

func datadetails() helpers.DataDetails {
	return helpers.DataDetails{
		Asin:                  "B08N5WRWNW",
		ProductTitle:          "Wireless Noise Cancelling Headphones",
		ProductPrice:          "$199.99",
		ProductOriginalPrice:  "$249.99",
		Currency:              "USD",
		Country:               "US",
		ProductByline:         "Brand X",
		ProductBylineLink:     "https://www.amazon.com/brandx",
		ProductStarRating:     "4.5",
		ProductNumRatings:     15234,
		ProductURL:            "https://www.amazon.com/dp/B08N5WRWNW",
		ProductPhoto:          "https://images.amazon.com/B08N5WRWNW.jpg",
		ProductNumOffers:      10,
		ProductAvailability:   "In Stock",
		IsBestSeller:          true,
		IsAmazonChoice:        false,
		IsPrime:               true,
		ClimatePledgeFriendly: false,
		SalesVolume:           "High",
		ProductDescription:    "Premium wireless headphones with active noise cancellation.",
		VideoThumbnail:        "https://images.amazon.com/B08N5WRWNW-video.jpg",
		HasVideo:              true,
		CustomersSay:          "Great battery life and sound quality!",
		Delivery:              "Free delivery",
		PrimaryDeliveryTime:   "2-3 business days",
		HasAPlus:              true,
		HasBrandStory:         false,
	}
}

func randomLeaderboardentry() leaderboard.LeaderboardEntry {
	return leaderboard.LeaderboardEntry{
		ProductId: util.RandomString(7),
		Score:     float64(util.RandomInt(0, 9)),
	}
}

func randomleaderboard() []*review.LeaderBoard {
	var lb []*review.LeaderBoard
	entries := randomLeaderBoardEntries()

	for _, entry := range entries {
		data := &review.LeaderBoard{
			ProductId: entry.ProductId,
			Score:     float32(entry.Score),
			ProductDetails: &review.ProductDetails{
				Asin:                 entry.ProductId,
				ProductTitle:         util.RandomString(7),
				ProductPrice:         util.RandomString(4),
				ProductOriginalPrice: util.RandomString(4),
				Currency:             util.RandomCurrency(),
				Country:              util.RandomCountry(),
				ProductUrl:           util.RandomString(8),
				ProductPhoto:         util.RandomString(7),
				ProductAvailability:  util.RandomString(6),
			},
		}

		lb = append(lb, data)
	}

	return lb
}
