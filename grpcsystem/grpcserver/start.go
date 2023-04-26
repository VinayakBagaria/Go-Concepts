package grpcserver

import (
	"context"
	"log"
	"net"

	pb "go-concepts/grpcsystem"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	log.Printf("Received message from client: %s", in.Name)
	return &pb.MessageReply{Message: "Hello from the Server!"}, nil
}

func DoWork() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
