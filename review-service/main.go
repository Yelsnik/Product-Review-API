package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"review-service/clients"
	db "review-service/db/sqlc"
	"review-service/gapi"
	"review-service/helpers"
	"review-service/leaderboard"
	"review-service/review"
	"review-service/util"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	// connect to redis
	opts, err := redis.ParseURL(config.Redis)
	if err != nil {
		log.Fatal("could not connect to redis", err)
	}

	redisClient := redis.NewClient(opts)
	pong, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("connected to redis successfully", pong)

	// helpers
	h := helpers.NewHelper(config)

	// create grpc conn
	nlpconn, err := grpc.NewClient("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to nlp service: %v", err)
	}
	defer nlpconn.Close()

	// initialize the client
	client := clients.NewClientgRPC(nlpconn)

	leaderboard := leaderboard.NewLeaderBoardClient(redisClient)

	store := db.NewStore(conn)

	grpcServer(config, store, h, client, leaderboard)

}

func grpcServer(config util.Config, store db.Store, helpers helpers.Helpers, client clients.Client, leaderboard leaderboard.Leaderboard) {
	// create a new server
	server, err := gapi.NewServer(config, store, helpers, client, leaderboard)
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
