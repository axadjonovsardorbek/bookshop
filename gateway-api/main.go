package main

import (
	"api/api"
	"api/api/handler"
	"api/config"
	pb "api/genproto/books"
	"api/services"
	"fmt"
	"log"
)

func main() {
	cfg := config.Load()
	conn := services.ClientConn()

	bookClient := pb.NewBooksServiceClient(conn.Book)
	authorsClient := pb.NewAuthorsServiceClient(conn.Book)
	publisherClient := pb.NewPublishersServiceClient(conn.Book)
	transleterClient := pb.NewTranslatorsServiceClient(conn.Book)

	router := api.NewGin(&handler.Handler{
		Book:       bookClient,
		Author:     authorsClient,
		Publisher:  publisherClient,
		Translater: transleterClient,
	})

	err := router.Run(fmt.Sprintf("%s:%d", cfg.GATEWAYHOST, cfg.GATEWAYPORT))
	if err != nil {
		log.Fatal(err)
	}
}
