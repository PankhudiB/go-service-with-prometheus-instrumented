package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service/monitoring"
	"go-service/monitoring/metrics"
	"go-service/request"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(monitoring.CaptureIncomingMetrics)

	r.POST("/hello", handleRequest)
	r.GET("/metrics", monitoring.PrometheusHandler)

	monitoring.RegisterCustomMetrics(metrics.InboundRequestMetrics)

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
