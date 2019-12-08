package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/maykonlf/grpc-gateway-sample/internal/server"
	v1 "github.com/maykonlf/grpc-gateway-sample/pkg/api/v1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net/http"
	"strings"
)

func main() {
	serverCert, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		log.Fatalln("failed to create server cert", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(serverCert))
	v1.RegisterPingPongServer(grpcServer, server.NewMyServer())

	clientCert, err := credentials.NewClientTLSFromFile("server.crt", "")
	if err != nil {
		log.Fatalln("failed to create client cert", err)
	}

	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:8080",
		grpc.WithTransportCredentials(clientCert))
	if err != nil {
		log.Fatalln("failed to dial to grpc server", err)
	}

	router := runtime.NewServeMux()
	if err = v1.RegisterPingPongHandler(context.Background(), router, conn); err != nil {
		log.Fatalln("failed to register gateway", err)
	}

	log.Fatalln(http.ListenAndServeTLS(":8080", "server.crt", "server.key", httpGrpcRouter(grpcServer, router)))
}

func httpGrpcRouter(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	})
}
