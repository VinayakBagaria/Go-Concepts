package grpcclient

import (
	"log"
	"net/http"

	pb "go-concepts/grpcsystem"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type todoTask struct {
	Name        string
	Description string
	Done        bool
}

func DoWork() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("didn't connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewTodoServiceClient(conn)

	r := gin.Default()

	r.POST("/task", func(ctx *gin.Context) {
		var todo todoTask
		if err := ctx.BindJSON(&todo); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := c.CreateTodo(ctx, &pb.NewTodo{Name: todo.Name, Description: todo.Description, Done: todo.Done})
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"id": res.GetId()})
		}
	})

	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
