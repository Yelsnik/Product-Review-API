package main

import (
	"context"
	"log"
	"net"
	db "review-service/db/sqlc"
	"review-service/gapi"
	"review-service/helpers"
	"review-service/review"
	"review-service/util"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// load config
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// connect to database
	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("could not connect", err)
	}

	// helpers
	h := helpers.NewHelper(config)

	store := db.NewStore(conn)

	grpcServer(config, store, h)

}

func grpcServer(config util.Config, store db.Store, helpers helpers.Helpers) {
	// create a new server
	server, err := gapi.NewServer(config, store, helpers)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	// // create a new grpc server
	grpcServer := grpc.NewServer()
	review.RegisterReviewServer(grpcServer, server)
	review.RegisterProductServer(grpcServer, server)
	reflection.Register(grpcServer)

	// start the server to listen to grpc
	// requests on a specific port
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("starting gRPC server at %s ...", listener.Addr().String())

	// start the server
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server:", err)
	}
}
