package handler

import (
	book "api/genproto/books"
	"api/kafka"
)

type Handler struct {
	Book book.BooksServiceClient
	Author book.AuthorsServiceClient
	Publisher book.PublishersServiceClient
	Translater book.TranslatorsServiceClient
	Kaf *kafka.KafkaProducer
	
	book.UnimplementedAuthorsServiceServer
	book.UnimplementedBooksServiceServer
	book.UnimplementedTranslatorsServiceServer
	book.UnimplementedPublishersServiceServer
}
