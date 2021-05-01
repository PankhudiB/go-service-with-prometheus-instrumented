package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"go-service/constants"
)

var InboundRequestMetrics = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "incoming_requests_metric",
		Help: "Metric for incoming HTTP requests",
	},
	[]string{
		constants.MONITORING_METHOD_KEY,
		constants.MONITORING_HOST_KEY,
		constants.MONITORING_PATH_KEY,
		constants.MONITORING_STATUS_CODE_KEY,
	},
)
