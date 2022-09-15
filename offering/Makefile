DEV_VERSION=0.0.0-dev
ENV=env GOOS=linux
export COURIER_MANAGEMENT_WORKING_DIR := ${GOPATH}/src/gitlab.artin.ai/backend/courier-management
export COURIER_MANAGEMENT_CONFIG_FILE := config

export TILE38_IP := 9851

.PHONY: test start test-storage

test:
	go test ./...

start:
	go run ../. start offering

test-storage:
	docker run -p ${TILE38_IP}:9851 -d tile38/tile38
	go test ./storage/... -v
	docker rm $$(docker stop $$(docker ps -a -q --filter ancestor=tile38/tile38 --format="{{.ID}}"))
