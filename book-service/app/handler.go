package app

import (
	"context"
	"log"

	pb "book/genproto/books"
	"book/service"

	"google.golang.org/protobuf/encoding/protojson"
)

type KafkaHandler struct {
	books       *service.BooksService
	authors     *service.AuthorsService
	publishers  *service.PublishersService
	translators *service.TranslatorsService
}

func (h *KafkaHandler) BooksCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.BooksRes
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.books.Create(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create books via Kafka: %v", err)
			return
		}
		log.Printf("Created books: %+v", res)
	}
}
func (h *KafkaHandler) BooksUpdateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.BooksUpdateReq
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.books.Update(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update books via Kafka: %v", err)
			return
		}
		log.Printf("Updated books: %+v", res)
	}
}
func (h *KafkaHandler) AuthorsCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.AuthorsRes
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.authors.Create(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create authors via Kafka: %v", err)
			return
		}
		log.Printf("Created authors: %+v", res)
	}
}
func (h *KafkaHandler) AuthorsUpdateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.AuthorsUpdateReq
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.authors.Update(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update authors via Kafka: %v", err)
			return
		}
		log.Printf("Updated authors: %+v", res)
	}
}
func (h *KafkaHandler) PublishersCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.PublishersRes
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.publishers.Create(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create publishers via Kafka: %v", err)
			return
		}
		log.Printf("Created publishers: %+v", res)
	}
}
func (h *KafkaHandler) PublishersUpdateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.PublishersUpdateReq
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.publishers.Update(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update publishers via Kafka: %v", err)
			return
		}
		log.Printf("Updated publishers: %+v", res)
	}
}
func (h *KafkaHandler) TranslatorsCreateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.TranslatorsRes
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.translators.Create(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot create translators via Kafka: %v", err)
			return
		}
		log.Printf("Created translators: %+v", res)
	}
}
func (h *KafkaHandler) TranslatorsUpdateHandler() func(message []byte) {
	return func(message []byte) {
		var req pb.TranslatorsUpdateReq
		if err := protojson.Unmarshal(message, &req); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := h.translators.Update(context.Background(), &req)
		if err != nil {
			log.Printf("Cannot update translators via Kafka: %v", err)
			return
		}
		log.Printf("Updated translators: %+v", res)
	}
}
