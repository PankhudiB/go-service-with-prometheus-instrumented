package monitoring

import (
	"github.com/gin-gonic/gin"
	"go-service/monitoring/metrics"
	"strconv"
)

func CaptureIncomingMetrics(ctx *gin.Context) {
	ctx.Next()
	method := ctx.Request.Method
	host := ctx.Request.Host
	path := ctx.FullPath()
	statusCode := strconv.Itoa(ctx.Writer.Status())
	metrics.InboundRequestMetrics.WithLabelValues(method, host, path, statusCode).Inc()
}
