package leaderboard

import (
	"context"
	"review-service/helpers"
	"review-service/review"
	"time"
)

type ProductDetails struct {
	Asin                 string
	ProductTitle         string
	ProductPrice         string
	ProductOriginalPrice string
	Currency             string
	Country              string
	ProductUrl           string
	ProductPhoto         string
	ProductAvailability  string
}

func (l *LeaderboardClient) GetProductdetails(ctx context.Context, leaderboard []LeaderboardEntry, helper helpers.Helpers) (*review.GetTop10ProductsResponse, error) {

	var cachedProducts []ProductDetails
	var uncachedProductIDS []string

	for _, v := range leaderboard {
		cached, err := l.redis.HGetAll(ctx, v.ProductId).Result()
		if err == nil && len(cached) > 0 {
			// Use cached details
			cachedProducts = append(cachedProducts, ProductDetails{
				Asin:                 v.ProductId,
				ProductTitle:         cached["product_title"],
				ProductPrice:         cached["product_price"],
				ProductOriginalPrice: cached["product_original_price"],
				Currency:             cached["currency"],
				Country:              cached["country"],
				ProductUrl:           cached["product_url"],
				ProductPhoto:         cached["product_photo"],
				ProductAvailability:  cached["product_availability"],
			})
		} else {
			uncachedProductIDS = append(uncachedProductIDS, v.ProductId)
		}
	}

	resultchan := make(chan ResponseDetails, len(uncachedProductIDS))
	errs := make(chan error, len(uncachedProductIDS))
	for _, productID := range uncachedProductIDS {
		go func(id string) {
			details, err := helper.GetAmazonProductDetails(productID, "US") // Your API client
			resultchan <- ResponseDetails{
				Status:     details.Status,
				RequestId:  details.RequestId,
				Parameters: details.Parameters,
				Data:       DataDetails(details.Data),
			}
			errs <- err
		}(productID)
	}

	var fetchedProducts []ProductDetails
	for i := 0; i < len(uncachedProductIDS); i++ {
		result := <-resultchan
		err := <-errs
		if err != nil {

		}
		l.redis.HSet(ctx, result.Data.Asin, ProductDetails{
			Asin:                 result.Data.Asin,
			ProductTitle:         result.Data.ProductTitle,
			ProductPrice:         result.Data.ProductPrice,
			ProductOriginalPrice: result.Data.ProductOriginalPrice,
			Currency:             result.Data.Currency,
			Country:              result.Data.Country,
			ProductUrl:           result.Data.ProductURL,
			ProductPhoto:         result.Data.ProductPhoto,
			ProductAvailability:  result.Data.ProductAvailability,
		})
		l.redis.Expire(ctx, result.Data.Asin, 24*time.Hour)

		fetchedProducts = append(fetchedProducts, ProductDetails{
			Asin:                 result.Data.Asin,
			ProductTitle:         result.Data.ProductTitle,
			ProductPrice:         result.Data.ProductPrice,
			ProductOriginalPrice: result.Data.ProductOriginalPrice,
			Currency:             result.Data.Currency,
			Country:              result.Data.Country,
			ProductUrl:           result.Data.ProductURL,
			ProductPhoto:         result.Data.ProductPhoto,
			ProductAvailability:  result.Data.ProductAvailability,
		})
	}
	close(resultchan)

	allProducts := append(cachedProducts, fetchedProducts...)
	lb := make([]*review.LeaderBoard, len(leaderboard))

	for i, v := range leaderboard {
		proudctId := v.ProductId

		for _, vp := range allProducts {
			if vp.Asin == proudctId {
				lb[i] = &review.LeaderBoard{
					ProductId: proudctId,
					Score:     float32(v.Score),
					ProductDetails: &review.ProductDetails{
						Asin:                 vp.Asin,
						ProductTitle:         vp.ProductTitle,
						ProductPrice:         vp.ProductPrice,
						ProductOriginalPrice: vp.ProductOriginalPrice,
						Currency:             vp.Currency,
						Country:              vp.Country,
						ProductUrl:           vp.ProductUrl,
						ProductPhoto:         vp.ProductPhoto,
						ProductAvailability:  vp.ProductAvailability,
					},
				}
				break
			}
		}
	}

	res := &review.GetTop10ProductsResponse{
		Leaderboard: lb,
	}

	return res, nil
}
