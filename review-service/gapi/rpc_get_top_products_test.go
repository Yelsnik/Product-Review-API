package gapi

import (
	"context"
	mockcl "review-service/clients/mock"
	mockdb "review-service/db/mock"
	mockhlp "review-service/helpers/mock"
	"review-service/leaderboard"
	mocklb "review-service/leaderboard/mock"
	"review-service/review"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func randomLeaderBoardEntries() []leaderboard.LeaderboardEntry {
	var lb []leaderboard.LeaderboardEntry
	for i := 0; i < 3; i++ {
		newLb := randomLeaderboardentry()

		lb = append(lb, newLb)
	}

	return lb
}

func randomTop10Products(lb []*review.LeaderBoard) *review.GetTop10ProductsResponse {
	return &review.GetTop10ProductsResponse{
		Leaderboard: lb,
	}
}

func TestGetTopProductAPI(t *testing.T) {
	leaderboardentries := randomLeaderBoardEntries()
	leaderb := randomleaderboard()
	top10products := randomTop10Products(leaderb)

	testCases := []struct {
		name          string
		req           *review.GetTop10ProductsRequest
		buildStubs    func(helper *mockhlp.MockHelpers, lbd *mocklb.MockLeaderboard)
		checkResponse func(t *testing.T, res *review.GetTop10ProductsResponse, err error)
	}{
		{
			name: "OK",
			req:  &review.GetTop10ProductsRequest{},
			buildStubs: func(helper *mockhlp.MockHelpers, lbd *mocklb.MockLeaderboard) {
				ten := int64(10)
				lbd.EXPECT().
					GetTopProducts(gomock.Any(), gomock.Eq(ten)).
					Times(1).
					Return(leaderboardentries, nil)

				lbd.EXPECT().
					GetProductdetails(gomock.Any(), leaderboardentries, helper).
					Times(1).
					Return(top10products, nil)
			},
			checkResponse: func(t *testing.T, res *review.GetTop10ProductsResponse, err error) {
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
			tc.buildStubs(helpers, lb)

			// initialize new test server
			server := newTestServer(t, store, helpers, client, lb)

			// call the grpc function
			res, err := server.GetTopProducts(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
