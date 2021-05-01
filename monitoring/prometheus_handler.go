package monitoring

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func PrometheusHandler(ctx *gin.Context) {
	handler := promhttp.Handler()
	handler.ServeHTTP(ctx.Writer, ctx.Request)
	return
}