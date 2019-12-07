CERT := $(shell which openssl)

all: certs

certs:
	openssl genrsa -out certs/server.key 2048
	openssl req -new -x509 -sha256 -key certs/server.key -out certs/server.crt -subj "/CN=localhost"

.PHONY: certs
