DEV_VERSION=0.0.0-dev
ENV=env GOOS=linux

export COURIER_MANAGEMENT_WORKING_DIR := ${GOPATH}/src/gitlab.artin.ai/backend/courier-management
export COURIER_MANAGEMENT_CONFIG_FILE := config

export GRPC_DIR := ${COURIER_MANAGEMENT_WORKING_DIR}/grpc

export GRPC_VALIDATE_DIR := ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate
export GOOGLE_API_DIR := ${GOPATH}/src/github.com/googleapis/googleapis
export GOOGLE_PROTOBUF_DIR := ${GOPATH}/src/github.com/gogo/protobuf

.PHONY: generate generate-finance generate-offering generate-delivery clean clean-go-finance clean-go-offering clean-go-delivery grpcui grpcui-finance grpcui-offering grpcui-delivery protod clean-go-common generate-common

generate: export GENERATED_GO_DIR=${GRPC_DIR}/${SVC}/go
#generate: export GENERATED_DART_DIR=${GRPC_DIR}/${SVC}/dart
generate: export PROTO_DIR=${GRPC_DIR}/${SVC}/proto
#generate: export DART_PROTO_DIR=${GRPC_DIR}/${SVC}/protod
generate: clean
generate: protod
	if ! cd ${GRPC_VALIDATE_DIR}; \
		then git clone --depth 1 --single-branch --branch v0.5.0 https://github.com/envoyproxy/protoc-gen-validate.git ${GRPC_VALIDATE_DIR}; fi

	if ! cd ${GOOGLE_API_DIR}; \
		then git clone --depth 1 --single-branch https://github.com/googleapis/googleapis.git ${GOOGLE_API_DIR}; fi

	if ! cd ${GOOGLE_PROTOBUF_DIR}; \
		then git clone --depth 1 --single-branch https://github.com/gogo/protobuf.git ${GOOGLE_PROTOBUF_DIR}; fi

	protoc \
		-I $(PROTO_DIR) \
		-I ${GRPC_VALIDATE_DIR} \
		-I ${GOOGLE_API_DIR} \
		-I $(GRPC_DIR)/common/proto \
		--validate_out=lang=go:${GOPATH}/src \
		--go_out=${GOPATH}/src --go-grpc_out=${GOPATH}/src \
 		-Iproto $$(find $(PROTO_DIR) -iname "*.proto")
	git add $(GENERATED_GO_DIR)/*
#
#	protoc \
#    	-I $(DART_PROTO_DIR) \
#    	-I $(GOOGLE_PROTOBUF_DIR)/protobuf \
#    	-I ${GOOGLE_API_DIR} \
#    	-I $(GRPC_DIR)/common/protod \
#		-Iproto $$(find $(DART_PROTO_DIR) -iname "*.proto") \
#    	-Iproto $(GOOGLE_PROTOBUF_DIR)/protobuf/google/protobuf/*.proto \
#        --dart_out=$(GENERATED_DART_DIR)
#	git add $(GENERATED_DART_DIR)/*

clean: export GENERATED_GO_DIR=${GRPC_DIR}/${SVC}/go
#clean: export GENERATED_DART_DIR=${GRPC_DIR}/${SVC}/dart
clean:
	rm -f -r $(GENERATED_GO_DIR)/*.*
#	rm -f -r $(GENERATED_DART_DIR)/*.*

protod:
	cp -a ${GRPC_DIR}/offering/proto/* ${GRPC_DIR}/offering/protod/
	sed -i 's/\[(validate.rules).*/;/g' ${GRPC_DIR}/offering/protod/*.proto
	sed -i 's/import "validate\/validate.proto";/ /g' ${GRPC_DIR}/offering/protod/*.proto
	git add ${GRPC_DIR}/offering/protod/*

	cp -a ${GRPC_DIR}/delivery/proto/* ${GRPC_DIR}/delivery/protod/
	sed -i 's/\[(validate.rules).*/;/g' ${GRPC_DIR}/delivery/protod/*.proto
	sed -i 's/import "validate\/validate.proto";/ /g' ${GRPC_DIR}/delivery/protod/*.proto
	git add ${GRPC_DIR}/delivery/protod/*

	cp -a ${GRPC_DIR}/finance/proto/* ${GRPC_DIR}/finance/protod/
	sed -i 's/\[(validate.rules).*/;/g' ${GRPC_DIR}/finance/protod/*.proto
	sed -i 's/import "validate\/validate.proto";/ /g' ${GRPC_DIR}/finance/protod/*.proto
	git add ${GRPC_DIR}/finance/protod/*

	cp -a ${GRPC_DIR}/pricing/proto/* ${GRPC_DIR}/pricing/protod/
	sed -i 's/\[(validate.rules).*/;/g' ${GRPC_DIR}/pricing/protod/*.proto
	sed -i 's/import "validate\/validate.proto";/ /g' ${GRPC_DIR}/pricing/protod/*.proto
	git add ${GRPC_DIR}/pricing/protod/*

	cp -a ${GRPC_DIR}/common/proto/* ${GRPC_DIR}/common/protod/
	sed -i 's/\[(validate.rules).*/;/g' ${GRPC_DIR}/common/protod/*.proto
	sed -i 's/import "validate\/validate.proto";/ /g' ${GRPC_DIR}/common/protod/*.proto
	git add ${GRPC_DIR}/common/protod/*

generate-finance: export SVC=finance
generate-finance: generate

generate-pricing: export SVC=pricing
generate-pricing: generate

generate-delivery: export SVC=delivery
generate-delivery: generate

generate-offering: export SVC=offering
generate-offering: generate

generate-common: export SVC=common
generate-common: generate



clean-go-finance: export SVC=finance
clean-go-finance: clean

clean-go-delivery: export SVC=delivery
clean-go-delivery: clean

clean-go-offering: export SVC=offering
clean-go-offering: clean

clean-go-pricing: export SVC=pricing
clean-go-pricing: clean

clean-go-common: export SVC=common
clean-go-common: clean

grpcui: export COMMON_DIR=$(GRPC_DIR)/common/proto
grpcui:
	grpcui -import-path $(GRPC_VALIDATE_DIR) \
		-import-path $(GOOGLE_API_DIR) \
		-import-path $(PROTO_DIR) \
		-import-path $(COMMON_DIR) \
		-proto $(SERVICE) \
		-port $(GRPCUI_PORT) \
		-vv \
		-plaintext $(HOST):$(PORT)

grpcui-finance-gke: export HOST="35.239.100.128"
grpcui-finance-gke: grpcui-finance

grpcui-finance-okd: export HOST="185.97.117.71"
grpcui-finance-okd: grpcui-finance

grpcui-finance-local: export HOST="localhost"
grpcui-finance-local: grpcui-finance

grpcui-finance: export SERVICE="finance.proto"
grpcui-finance: export PORT="50002"
grpcui-finance: export PROTO_DIR=${GRPC_DIR}/finance/proto
grpcui-finance: export GRPCUI_PORT="7090"
grpcui-finance: grpcui

grpcui-delivery-gke: export HOST="35.238.133.12"
grpcui-delivery-gke: grpcui-delivery

grpcui-delivery-okd: export HOST="185.97.117.71"
grpcui-delivery-okd: grpcui-delivery

grpcui-delivery-local: export HOST="localhost"
grpcui-delivery-local: grpcui-delivery

grpcui-delivery: export SERVICE="delivery.proto"
grpcui-delivery: export PORT="50001"
grpcui-delivery: export PROTO_DIR=${GRPC_DIR}/delivery/proto
grpcui-delivery: export GRPCUI_PORT="7071"
grpcui-delivery: grpcui

grpcui-offering-gke: export HOST="34.123.194.31"
grpcui-offering-gke: grpcui-offering

grpcui-offering-okd: export HOST="185.97.117.71"
grpcui-offering-okd: grpcui-offering

grpcui-offering-local: export HOST="localhost"
grpcui-offering-local: grpcui-offering

grpcui-offering: export SERVICE="offering.proto"
grpcui-offering: export PORT="50003"
grpcui-offering: export PROTO_DIR=${GRPC_DIR}/offering/proto
grpcui-offering: export GRPCUI_PORT="7072"
grpcui-offering: grpcui

grpcui-pricing: export SERVICE="pricing.proto"
grpcui-pricing: export PORT="50004"
grpcui-pricing: export PROTO_DIR=${GRPC_DIR}/pricing/proto
grpcui-pricing: export GRPCUI_PORT="7073"
grpcui-pricing: grpcui

grpcui-uaa-gke: export HOST="35.202.133.218"
grpcui-uaa-gke: grpcui-uaa

grpcui-uaa-okd: export HOST="185.97.117.71"
grpcui-uaa-okd: grpcui-uaa

grpcui-uaa-local: export HOST="localhost"
grpcui-uaa-local: grpcui-uaa

grpcui-uaa: export SERVICE="uaa.proto"
grpcui-uaa: export PORT="8086"
grpcui-uaa: export PROTO_DIR=${COURIER_MANAGEMENT_WORKING_DIR}/uaa/proto
grpcui-uaa: export GRPCUI_PORT="7074"
grpcui-uaa: grpcui

grpcui-party-gke: export HOST="34.66.6.44"
grpcui-party-gke: grpcui-party

grpcui-party-okd: export HOST="185.97.117.71"
grpcui-party-okd: grpcui-party

grpcui-party-local: export HOST="localhost"
grpcui-party-local: grpcui-party

grpcui-party: export SERVICE="party.proto"
grpcui-party: export PORT="8085"
grpcui-party: export PROTO_DIR=${COURIER_MANAGEMENT_WORKING_DIR}/party/proto
grpcui-party: export GRPCUI_PORT="7075"
grpcui-party: grpcui