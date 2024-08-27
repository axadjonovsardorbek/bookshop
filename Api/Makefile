CURRENT_DIR=$(shell pwd)

run:
	go run main.go

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

create-mig:
	migrate create -ext sql -dir migration -seq create_users

mig-insert:
	migrate create -ext sql -dir db/migrations -seq insert_table

SWAGGER := $(HOME)/go/bin/swag
SWAGGER_DOCS := docs
SWAGGER_INIT := $(SWAGGER) init -g ./api/router.go -o $(SWAGGER_DOCS)

swag-gen:
	@echo "Running: $(SWAGGER_INIT)"
	$(SWAGGER_INIT)


