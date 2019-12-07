package main

import (
	"context"
	"fmt"
	pingpong "github.com/maykonlf/grpc-gateway-sample/protos"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	clientCert, err := credentials.NewClientTLSFromFile("certs/server.crt", "")
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

	client := pingpong.NewPingPongClient(conn)

	for i := 0; i < 1000; i++ {
		response, err := client.Ping(context.Background(), &pingpong.PingRequest{
			Ping: fmt.Sprintf("%d", i),
		})
		if err != nil {
			log.Fatalln(err)
		}
		logrus.Info(map[string]string{"pong": response.Pong})
	}
}
