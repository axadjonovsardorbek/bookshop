package main

import (
	"auth/config"
	"auth/kafka"
	"auth/service"
	"auth/storage/postgres"
	"log"
	"net"

	pb "auth/genproto/auth"

	"google.golang.org/grpc"
)

func main() {

	cfg := config.Load()

	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatalln("Error while connecting to database", err)
	}
	log.Println("Database connected successfully! ")

	lis, err := net.Listen("tcp", cfg.HTTPPort)
	if err != nil {
		log.Fatal("Error while creating TCP listener", err)
	}
	defer lis.Close()

	kafka, err := kafka.NewKafkaProducer([]string{cfg.KafkaHost + cfg.HTTPPort})
	if err != nil {
		log.Fatal(err)
	}
	defer kafka.Close()

	server := grpc.NewServer()
	svc := service.NewAuthService(stg)
	pb.RegisterAuthServiceServer(server, svc)

	log.Println("Server listening at:", cfg.HTTPPort)
	if server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
