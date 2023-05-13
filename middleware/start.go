package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func DoWork() {
	engine := gin.New()

	engine.Use(traceRequest)

	engine.GET("query", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "got request",
		})
	})

	engine.Run(":8000")
}

func traceRequest(c *gin.Context) {
	beforeRequest(c)
	c.Next()
	afterRequest(c)
}

func beforeRequest(c *gin.Context) {
	c.Set("startTime", time.Now())
}

func afterRequest(c *gin.Context) {
	startTime, ok := c.Get("startTime")
	if !ok {
		fmt.Println("Error in getting time")
		return
	}

	duration := time.Since(startTime.(time.Time))
	fmt.Printf("Resolved in %v\n", duration)
}
