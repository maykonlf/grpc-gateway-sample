GRPC_GATEWAY_PATH := $(shell go list -m -f '{{.Dir}}' all | grep 'github.com/grpc-ecosystem/grpc-gateway')
GRPC_GOOGLE_APIS_PATH := $(GRPC_GATEWAY_PATH)/third_party/googleapis/

all: certs proto build

certs:
	openssl genrsa -out server.key 2048
ifdef cn
	openssl req -new -x509 -sha256 -key server.key -out server.crt -subj "/CN=$(cn)"
else
	openssl req -new -x509 -sha256 -key server.key -out server.crt -subj "/CN=localhost"
endif

proto:
	protoc -I api/proto/v1/ -I "$(GRPC_GOOGLE_APIS_PATH)" -I "$(GRPC_GATEWAY_PATH)" --go_out=plugins=grpc:pkg/api/v1/ --grpc-gateway_out=logtostderr=true:pkg/api/v1/ --swagger_out=logtostderr=true:api/swagger/v1/ pingpong.proto

build:
	CGO_ENABLED=0
	GOOS=linux
	GOARCH=amd64
	go get ./...
	go build -o server cmd/server/main.go
	go build -o client cmd/client/main.go

.PHONY: certs proto build
