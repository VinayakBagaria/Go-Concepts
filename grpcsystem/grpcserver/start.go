package grpcserver

import (
	"context"
	"log"
	"net"

	pb "go-concepts/grpcsystem"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTodoServiceServer
}

func (s *server) CreateTodo(ctx context.Context, in *pb.NewTodo) (*pb.Todo, error) {
	log.Printf("Received message from client: %s", in.GetName())
	return &pb.Todo{
		Name:        in.GetName(),
		Description: in.GetDescription(),
		Done:        in.GetDone(),
		Id:          uuid.New().String(),
	}, nil
}

func DoWork() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
