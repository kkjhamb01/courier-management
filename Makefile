DEV_VERSION=0.0.0-dev
ENV=env GOOS=linux
export COURIER_MANAGEMENT_WORKING_DIR := ${GOPATH}/src/courier-management
export COURIER_MANAGEMENT_CONFIG_FILE := config

export COURIER_MANAGEMENT_DOCKER_IMG_NAME := artin/courier_management

export TILE38_IP := 9851

.PHONY: build test start-offering start-finance start-delivery clean start-all test-storage build-docker-image start-docker-container

build:
	go build -o courier_management .

test:
	go test ./...

start-offering: export app-port = "50003"
start-offering: export app = "offering"
start-offering:
	make start-docker-container

start-delivery: export app-port = "50001"
start-delivery: export app = "delivery"
start-delivery:
	make start-docker-container

start-finance: export app-port = "50002"
start-finance: export app = "finance"
start-finance:
	make start-docker-container

clean:
	rm -f courier_management

# We are not going to use this, the target should probably get removed
start-all:
	go run . start

test-storage:
	make -C offering test-storage
	# TODO add storage test for other services

build-docker-image:
	DOCKER_BUILDKIT=1 docker build -t ${COURIER_MANAGEMENT_DOCKER_IMG_NAME} .

start-docker-container:
	make build-docker-image
	docker run -p ${app-port}:${app-port} --env-file .env --env SERVICE_NAME=${app} ${COURIER_MANAGEMENT_DOCKER_IMG_NAME}

generate-all-services-grpc-go:
	$(MAKE) -C ./grpc generate-go-offering
	$(MAKE) -C ./grpc generate-go-finance
	$(MAKE) -C ./grpc generate-go-delivery
