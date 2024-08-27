package services

import (
	"api/config"
	"fmt"
	"log"
	_ "api/docs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	Book *grpc.ClientConn
	Auth *grpc.ClientConn
}

func ClientConn() *Clients {
	cfg := config.Load()
	book, err := grpc.NewClient(fmt.Sprintf("%s:%d", cfg.BOOKHOST, cfg.BOOKPORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while connecting to book", err)
		return nil
	}
	auth, err := grpc.NewClient(fmt.Sprintf("%s:%d", cfg.AUTHHOST, cfg.AUTHPORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while connecting to auth", err)
		return nil
	}
	return &Clients{
		Book: book,
		Auth: auth,
	}
}
