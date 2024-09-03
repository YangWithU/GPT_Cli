package main

import (
	pb "GPT_cli/grpc/gen"
	"GPT_cli/requests"
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedStringServiceServer
	processedString string
}

func SendContent(msg string) string {
	c, err := requests.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	res, err := c.SimpleSend(ctx, msg)
	if err != nil {
		log.Fatal(err)
	}

	a, _ := json.MarshalIndent(res, "", "  ")
	return string(a)
}

func (s *server) SendString(ctx context.Context, msg *pb.StringMessage) (*pb.EmptyMessage, error) {
	// here, we send the instructions to GPT
	s.processedString = SendContent(msg.Content)
	log.Printf("ASR message sent: %s", msg.Content)
	return &pb.EmptyMessage{}, nil
}

func (s *server) ReceiveString(ctx context.Context, empty *pb.EmptyMessage) (*pb.StringMessage, error) {
	return &pb.StringMessage{Content: s.processedString}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStringServiceServer(s, &server{})

	log.Printf("Server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
