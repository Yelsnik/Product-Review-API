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

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	// run db migrations

	// connect to redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword, // no password set
		DB:       0,                    // use default DB
	})

	//
	// opts, err := redis.ParseURL(config.Redis)
	// if err != nil {
	// 	log.Fatal("could not connect to redis: ", err)
	// }

	// redisClient := redis.NewClient(opts)
	pong, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("connected to redis successfully", pong)

	// helpers
	h := helpers.NewHelper(config)

	// create grpc conn
	nlpconn, err := grpc.NewClient("switchyard.proxy.rlwy.net:21251", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to nlp service: %v", err)
	}
	defer nlpconn.Close()

	// initialize the client
	client := clients.NewClientgRPC(nlpconn)

	leaderboard := leaderboard.NewLeaderBoardClient(redisClient)

	runDBmigration(config.MigrationURL, config.DBSource)
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

func runDBmigration(migrationURL, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

	log.Println("db migrated succesfully!")
}
