package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-service/monitoring"
	"go-service/monitoring/metrics"
	"go-service/request"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()

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
		captureMetrics(ctx)
		return
	}
	fmt.Println("Message :", req.Message)
	ctx.Status(http.StatusOK)
	captureMetrics(ctx)
	return
}

func captureMetrics(ctx *gin.Context) {
	method := ctx.Request.Method
	host := ctx.Request.Host
	path := ctx.FullPath()
	statusCode := strconv.Itoa(ctx.Writer.Status())
	metrics.InboundRequestMetrics.WithLabelValues(method, host, path, statusCode)
}
