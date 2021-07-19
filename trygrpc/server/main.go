package main

import (
	"context"
	"log"
	"net"
	"playground/trygrpc/chat"

	"google.golang.org/grpc"
)

type server struct {
	chat.UnimplementedChatServiceServer
}

func (s *server) SayHello(ctx context.Context, in *chat.HelloRequest) (*chat.HelloResponse, error) {
	log.Printf("receive message body from client %s", in.Body)
	return &chat.HelloResponse{Body: "Hello from the server!"}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %+v", err)
	}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
