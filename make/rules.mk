.PHONY: generate
generate:
	echo "generage grpc proto"
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