package main

import (
	"log"
	pb "playground/grpcdemo/chat"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &pb.HelloRequest{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
