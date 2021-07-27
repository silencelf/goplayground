package main

import (
	"context"
	"io"
	"log"
	"net"
	pb "playground/grpcdemo/chat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("receive message body from client %s", in.Body)
	return &pb.HelloResponse{Body: "Hello from the server!"}, nil
}

func (s *server) Average(stream pb.ChatService_AverageServer) error {
	var values []float32
	average := float32(0)
	for {
		v, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageMessage{
				Value: average,
			})
		}
		if err != nil {
			return err
		}
		// this is straight forward but not efficient,
		// I'll leave it here for now
		values = append(values, v.Value)
		sum := float32(0)
		for i := 0; i < len(values); i++ {
			sum += values[i]
		}
		average = sum / float32(len(values))
	}
}

func (s *server) Max(stream pb.ChatService_MaxServer) error {
	// is this thread safe?
	values := []int32{1}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("failed to recieve a value: %v", err)
			}
			log.Printf("got message %v", in)
			values = append(values, in.Value)
			log.Printf("values: %v", values)
		}
	}()

	for _, v := range values {
		if err := stream.Send(&pb.MaxMessage{Value: v}); err != nil {
			log.Fatalf("failed to send a value: %v", err)
		}
	}
	<-waitc
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %+v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)
	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
