package grpcclient

import (
	"context"
	"fmt"
	"log"

	pb "go-concepts/grpcsystem"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DoWork() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("didn't connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &pb.MessageRequest{Name: "Hello from client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("1. Response from server: %s", response.GetMessage())

	fmt.Println("--------------------------------------------------------------------------------")

	response, err = c.SayHello(context.Background(), &pb.MessageRequest{Name: "New Hello from client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("2. Response from server: %s", response.GetMessage())
}
