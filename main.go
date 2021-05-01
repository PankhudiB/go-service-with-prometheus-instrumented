package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service/request"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/hello", func(ctx *gin.Context) {
		handleRequest(ctx)
	})

	err := http.ListenAndServe(":8090", r)
	if err != nil {
		fmt.Println("Could not start service", err)
	}
}

func handleRequest(ctx *gin.Context) {
	fmt.Println("------------------ Welcome ------------------")
	var req request.HelloRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Error reading hello request", err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	fmt.Println("Message :", req.Message)
	ctx.Status(http.StatusOK)
	return
}
