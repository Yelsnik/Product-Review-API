package gapi

import (
	"review-service/clients"
	db "review-service/db/sqlc"
	"review-service/helpers"
	"review-service/leaderboard"
	"review-service/review"
	"review-service/util"
)

// serves gRPC requests for our banking service
type Server struct {
	review.UnimplementedReviewServer
	review.UnimplementedProductServer
	config      util.Config
	store       db.Store
	helpers     helpers.Helpers
	client      clients.Client
	leaderboard leaderboard.Leaderboard
}

// creates a new gRPC server
func NewServer(config util.Config, store db.Store, helpers helpers.Helpers, client clients.Client, leaderboard leaderboard.Leaderboard) (*Server, error) {

	server := &Server{
		config:      config,
		store:       store,
		helpers:     helpers,
		client:      client,
		leaderboard: leaderboard,
	}

	return server, nil
}
