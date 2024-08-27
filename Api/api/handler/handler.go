package handler

import (
	pb "api/genproto/auth"
	"api/kafka"

	"github.com/go-redis/redis/v8"
)

type Handler struct {
	Auth     pb.AuthServiceClient
	Rdb      *redis.Client
	Producer kafka.KafkaProducer
}

func NewHandler(auth pb.AuthServiceClient, rdb redis.Client, kafka kafka.KafkaProducer) *Handler {
	return &Handler{
		Auth:     auth,
		Rdb:      &rdb,
		Producer: kafka,
	}
}
