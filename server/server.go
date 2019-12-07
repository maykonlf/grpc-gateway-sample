package main

import (
	"context"
	pingpong "github.com/maykonlf/grpc-gateway-sample/protos"
)

type MyServer struct {}

func (m *MyServer) Ping(ctx context.Context, r *pingpong.PingRequest) (*pingpong.PingResponse, error) {
	return &pingpong.PingResponse{Pong: r.Ping,}, nil
}

func NewMyServer() pingpong.PingPongServer {
	return &MyServer{}
}
