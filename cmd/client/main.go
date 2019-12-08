package main

import (
	"context"
	"fmt"
	v1 "github.com/maykonlf/grpc-gateway-sample/pkg/api/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
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
	defer conn.Close()

	client := v1.NewPingPongClient(conn)

	for i := 0; i < 1000; i++ {
		response, err := client.Ping(context.Background(), &v1.PingRequest{
			Ping: fmt.Sprintf("%d", i),
		})
		if err != nil {
			log.Fatalln(err)
		}
		logrus.Info(map[string]string{"pong": response.Pong})
	}
}
