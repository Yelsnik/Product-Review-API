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

	response, err := server.leaderboard.GetProductdetails(ctx, leaderboard, server.helpers)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get product details: %s", err)
	}

	return response, nil
}
