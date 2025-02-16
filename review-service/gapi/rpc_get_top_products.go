package gapi

import (
	"context"
	"review-service/review"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetTopProducts(ctx context.Context, req *review.GetTop10ProductsRequest) (*review.GetTop10ProductsResponse, error) {

	leaderboard, err := server.leaderboard.GetTopProducts(ctx, 10)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get products from leaderboard: %s", err)
	}

	var lb []*review.LeaderBoard

	for _, v := range leaderboard {
		entry := &review.LeaderBoard{
			ProductId: v.ProductId,
			Score:     float32(v.Score),
		}

		lb = append(lb, entry)
	}

	response := &review.GetTop10ProductsResponse{
		Leaderboard: lb,
	}

	return response, nil
}
