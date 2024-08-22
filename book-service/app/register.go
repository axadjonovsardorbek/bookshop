package app

import (
	"errors"

	"book/config"
	"book/kafka"
)

func Register(h *KafkaHandler, cfg *config.Config) error {

	brokers := []string{cfg.KAFKA_HOST + cfg.KAFKA_PORT}
	kcm := kafka.NewKafkaConsumerManager()

	if err := kcm.RegisterConsumer(brokers, "books-create", "books-create-id", h.BooksCreateHandler()); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			return errors.New("consumer for topic 'books-create' already exists")
		} else {
			return errors.New("error registering consumer:" + err.Error())
		}
	}
	if err := kcm.RegisterConsumer(brokers, "books-update", "books-update-id", h.BooksUpdateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'books-update' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "authors-create", "authors-create-id", h.AuthorsCreateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'authors-create' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "authors-update", "authors-update-id", h.AuthorsUpdateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'authors-update' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "publishers-create", "publishers-create-id", h.PublishersCreateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'publishers-create' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "publishers-update", "publishers-update-id", h.PublishersUpdateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'publishers-update' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "translators-create", "translators-create-id", h.TranslatorsCreateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'translators-create' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	if err := kcm.RegisterConsumer(brokers, "translators-update", "translators-update-id", h.TranslatorsUpdateHandler()); err!= nil {
		if err == kafka.ErrConsumerAlreadyExists {
            return errors.New("consumer for topic 'translators-update' already exists")
        } else {
            return errors.New("error registering consumer:" + err.Error())
        }
	}
	return nil

}
