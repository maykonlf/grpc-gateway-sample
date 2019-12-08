package server

import (
	"context"
	v1 "github.com/maykonlf/grpc-gateway-sample/pkg/api/v1"
)

type MyServer struct {}

func (m *MyServer) Ping(ctx context.Context, r *v1.PingRequest) (*v1.PingResponse, error) {
	return &v1.PingResponse{Pong: r.Ping,}, nil
}

func NewMyServer() v1.PingPongServer {
	return &MyServer{}
}
