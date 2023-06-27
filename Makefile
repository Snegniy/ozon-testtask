.PHONY: run
run:
	go run cmd/main.go

.PHONY: local
local:
	docker compose down --remove-orphans
	docker compose --profile localdb up --build

.PHONY: postgres
postgres:
	docker compose down --remove-orphans
	docker compose --profile postgres up --build

.PHONY: stop
stop:
	docker compose down --remove-orphans

.PHONY: test
test:
	go test -v -race -cover ./...

.PHONY: proto
proto:
			mkdir -p pkg/api
			protoc -I api/proto \
            --go_out=pkg/api --go_opt=paths=source_relative \
            --plugin=protoc-gen-go=bin/protoc-gen-go \
            --go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative \
            --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
            api/proto/shortlinks.proto

.DEFAULT_GOAL := local