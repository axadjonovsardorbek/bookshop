package app

import (
	"book/config"
	"book/kafka"
	"book/service"

	// "book/storage/mongo"
	"book/storage/postgres"
	"log"
	"net"

	bp "book/genproto/books"

	"google.golang.org/grpc"
)

func Run(cfg config.Config) {

	db, err := postgres.NewPostgresStorage(cfg)

	if err != nil {
		panic(err)
	}

	// mongoConn, err := mongo.NewMongoStorage(cfg)

	// if err != nil {
	// 	panic(err)
	// }

	listener, err := net.Listen("tcp", cfg.BOOK_SERVICE_PORT)

	if err != nil {
		log.Fatalf("Failed to listen tcp: %v", err)
	}

	kafka, err := kafka.NewKafkaProducer([]string{cfg.KAFKA_HOST + cfg.KAFKA_PORT})
	if err != nil {
		log.Fatal(err)
	}

	books_service := service.NewBooksService(db)
	authors_service := service.NewAuthorsService(db)
	publishers_service := service.NewPublishersService(db)
	translators_service := service.NewTranslatorsService(db)

	kafka_handler := &KafkaHandler{
		books:       books_service,
		authors:     authors_service,
		publishers:  publishers_service,
		translators: translators_service,
	}

	if err := Register(kafka_handler, &cfg); err != nil {
		log.Fatal("Error registering kafka handlers: ", err)
	}

	defer kafka.Close()

	s := grpc.NewServer()

	bp.RegisterAuthorsServiceServer(s, service.NewAuthorsService(db))
	bp.RegisterBooksServiceServer(s, service.NewBooksService(db))
	bp.RegisterPublishersServiceServer(s, service.NewPublishersService(db))
	bp.RegisterTranslatorsServiceServer(s, service.NewTranslatorsService(db))

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
