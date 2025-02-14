package gapi

import (
	db "review-service/db/sqlc"
	"review-service/helpers"
	"review-service/review"
	"review-service/util"
)

// serves gRPC requests for our banking service
type Server struct {
	review.UnimplementedReviewServer
	review.UnimplementedProductServer
	config util.Config
	store  db.Store
	helpers helpers.Helpers
}

// creates a new gRPC server
func NewServer(config util.Config, store db.Store, helpers helpers.Helpers) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
		helpers: helpers,
	}

	return server, nil
}
