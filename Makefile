.PHONY: run
run:
	go run cmd/main.go

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


post:
	curl --request POST --data '{"url" : "ozon"}' http://localhost:8080/
get:
	curl --request GET --data '{"url" : "Li0QUvKTcT"}' http://localhost:8080/

.DEFAULT_GOAL := run
