package grpc

//import (
//	pb "GPT_cli/grpc/gen"
//	"context"
//	"google.golang.org/grpc"
//	"log"
//	"net"
//	"strings"
//)
//
//type server struct {
//	pb.UnimplementedStringServiceServer
//	processedString string
//}
//
//func (s *server) SendString(ctx context.Context, msg *pb.StringMessage) (*pb.EmptyMessage, error) {
//	// Process the string (for example, convert to uppercase)
//	s.processedString = strings.ToUpper(msg.Content)
//	log.Printf("Received and processed string: %s", s.processedString)
//	return &pb.EmptyMessage{}, nil
//}
//
//func (s *server) ReceiveString(ctx context.Context, empty *pb.EmptyMessage) (*pb.StringMessage, error) {
//	return &pb.StringMessage{Content: s.processedString}, nil
//}
//
//func init() {
//	lis, err := net.Listen("tcp", ":50051")
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	s := grpc.NewServer()
//	pb.RegisterStringServiceServer(s, &server{})
//
//	log.Printf("Server listening on %v", lis.Addr())
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}
