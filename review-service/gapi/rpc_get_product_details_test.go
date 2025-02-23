package gapi

import (
	"context"
	mockcl "review-service/clients/mock"
	mockdb "review-service/db/mock"
	"review-service/helpers"
	mockhlp "review-service/helpers/mock"
	mocklb "review-service/leaderboard/mock"
	"review-service/review"
	"review-service/util"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func randomProduct() *helpers.ResponseDetails {
	data := datadetails()
	return &helpers.ResponseDetails{
		Status:     "success",
		RequestId:  util.RandomString(7),
		Parameters: "",
		Data:       data,
	}
}

func TestGetProductDetailsAPI(t *testing.T) {
	productdetails := randomProduct()

	testCases := []struct {
		name          string
		req           *review.GetProductDetailsRequest
		buildStubs    func(store *mockdb.MockStore, helper *mockhlp.MockHelpers)
		checkResponse func(t *testing.T, res *review.GetProductDetailsResponse, err error)
	}{
		{
			name: "OK",
			req: &review.GetProductDetailsRequest{
				Asin:    productdetails.Data.Asin,
				Country: productdetails.Data.Country,
			},
			buildStubs: func(store *mockdb.MockStore, helper *mockhlp.MockHelpers) {
				helper.EXPECT().
					GetAmazonProductDetails(gomock.Any(), gomock.Eq(productdetails.Data.Asin), gomock.Eq(productdetails.Data.Country)).
					Times(1).
					Return(productdetails, nil)
			},
			checkResponse: func(t *testing.T, res *review.GetProductDetailsResponse, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, res)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			// store controller
			storeCtrl := gomock.NewController(t)
			defer storeCtrl.Finish()

			store := mockdb.NewMockStore(storeCtrl)

			// client controller
			clientCtrl := gomock.NewController(t)
			defer clientCtrl.Finish()

			client := mockcl.NewMockClient(clientCtrl)

			// helper controller
			helperCtrl := gomock.NewController(t)
			defer helperCtrl.Finish()

			helpers := mockhlp.NewMockHelpers(helperCtrl)

			// leaderboard controller
			lbCtrl := gomock.NewController(t)
			defer lbCtrl.Finish()

			lb := mocklb.NewMockLeaderboard(lbCtrl)

			// build stubs
			tc.buildStubs(store, helpers)

			// initialize new test server
			server := newTestServer(t, store, helpers, client, lb)

			// call the grpc function
			res, err := server.GetProductDetails(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
