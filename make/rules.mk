.PHONY: generate
generate:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/storage/storage.proto
.PHONY: build
build: generate
	docker-compose --file ./deployments/docker-compose.yml build 

.PHONY: run
run: 
	docker-compose --file ./deployments/docker-compose.yml up

.PHONY: test
test: 
	go test ./...

.PHONY: check
check:
	golangci-lint run