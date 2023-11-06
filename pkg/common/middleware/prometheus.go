package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func NewPrometheusMiddleware(p *Prometheus) gin.HandlerFunc {
	if p != nil {
		return p.HandlerFunc()
	}

	return func(c *gin.Context) {
		c.Next()
	}
}

func NewPrometheus(cfg config.Monitoring, logger *LoggerAdapter) *Prometheus {
	if !cfg.PrometheusEnabled {
		logger.Logger.Info("Prometheus disabled", map[string]interface{}{"from": common.Prometheus.String()})
		return nil
	}

	p := &Prometheus{
		metricsList:    standardMetrics,
		replaceURLKeys: replaceURLKeys,
		logger:         logger,
	}

	// Register metrics with prefix
	p.registerMetrics(cfg.PrometheusAppName)

	return p
}

// HandlerFunc is the actual middleware, it's where the magic happens
func (p *Prometheus) HandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Don't log the /metrics calls
		if c.Request.URL.Path == "/metrics" {
			c.Next()
			return
		}

		// Start request
		start := time.Now()
		requestSize := getApproxRequestSize(c.Request)

		c.Next()

		// Get relevant info
		method := c.Request.Method                                   // e.g. GET
		status := strconv.Itoa(c.Writer.Status())                    // e.g. 200
		endpoint := p.replaceURLKeys(c)                              // e.g. /users/:user_id
		elapsed := float64(time.Since(start)) / float64(time.Second) // e.g. 0.0123 (seconds)
		responseSize := float64(c.Writer.Size())                     // e.g. 1234 (bytes)

		// Increment & Observe metrics
		p.totalRequests.WithLabelValues(status, endpoint, method).Inc()
		p.requestsDuration.WithLabelValues(status, endpoint, method).Observe(elapsed)
		p.requestsSize.Observe(float64(requestSize))
		p.responsesSize.Observe(responseSize)
	}
}

// Prometheus contains the metrics gathered by the instance and its path
type Prometheus struct {
	metricsList []*Metric

	totalRequests    *prometheus.CounterVec
	requestsDuration *prometheus.HistogramVec
	requestsSize     prometheus.Summary
	responsesSize    prometheus.Summary

	replaceURLKeys func(c *gin.Context) string

	logger *LoggerAdapter
}

// prometheus.Collector type (i.e. CounterVec, Summary, etc) of each metric
type Metric struct {
	MetricCollector prometheus.Collector // the type of the metric: counter_vec, gauge, etc
	ID              string
	Name            string
	Description     string
	Type            string
	Args            []string
}

// Available metrics are:
//
//	counter, counter_vec, gauge, gauge_vec,
//	histogram, histogram_vec, summary, summary_vec

var standardMetrics = []*Metric{
	metricTotalRequests,
	metricRequestsDuration,
	metricResponsesSize,
	metricRequestsSize,
}

var metricTotalRequests = &Metric{
	ID:          "totalRequests",
	Name:        "total_requests",
	Description: "Total number of HTTP Requests received, to which endpoints.",
	Type:        "counter_vec",
	Args:        []string{"status", "endpoint", "method"},
}

var metricRequestsDuration = &Metric{
	ID:          "requestsDuration",
	Name:        "requests_duration",
	Description: "HTTP Requests latencies in seconds.",
	Type:        "histogram_vec",
	Args:        []string{"status", "endpoint", "method"},
}

var metricRequestsSize = &Metric{
	ID:          "requestsSize",
	Name:        "requests_size",
	Description: "HTTP Requests sizes in bytes.",
	Type:        "summary",
}

var metricResponsesSize = &Metric{
	ID:          "responsesSize",
	Name:        "responses_size",
	Description: "HTTP Responses sizes in bytes.",
	Type:        "summary",
}

// NewMetric associates prometheus.Collector based on Metric.Type
func NewMetric(m *Metric, subsystem string) (metric prometheus.Collector) {
	switch m.Type {
	case "counter_vec":
		metric = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name:      m.Name,
				Subsystem: subsystem,
				Help:      m.Description,
			},
			m.Args,
		)
	case "counter":
		metric = prometheus.NewCounter(
			prometheus.CounterOpts{
				Name:      m.Name,
				Subsystem: subsystem,
				Help:      m.Description,
			},
		)
	case "gauge_vec":
		metric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name:      m.Name,
				Subsystem: subsystem,
				Help:      m.Description,
			},
			m.Args,
		)
	case "gauge":
		metric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name:      m.Name,
				Subsystem: subsystem,
				Help:      m.Description,
			},
		)
	case "histogram_vec":
		metric = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:      m.Name,
				Subsystem: subsystem,
				Help:      m.Description,
			},
			m.Args,
		)
	case "histogram":
		metric = prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Name:      m.Name,
				Subsystem: subsystem,
				Help:      m.Description,
			},
		)
	case "summary_vec":
		metric = prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Name:      m.Name,
				Subsystem: subsystem,
				Help:      m.Description,
			},
			m.Args,
		)
	case "summary":
		metric = prometheus.NewSummary(
			prometheus.SummaryOpts{
				Name:       m.Name,
				Subsystem:  subsystem,
				Help:       m.Description,
				Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.95: 0.005, 0.99: 0.001},
			},
		)
	}
	return metric
}

func (p *Prometheus) registerMetrics(subsystem string) {

	// For each metric create the appropiate Collector and register it
	for _, metricDefinition := range p.metricsList {
		metric := NewMetric(metricDefinition, subsystem)
		if err := prometheus.Register(metric); err != nil {
			p.logger.Logger.Error(
				err.Error(),
				map[string]interface{}{
					"error": fmt.Errorf("%s could not be registered in Prometheus", metricDefinition.Name),
					"from":  common.Prometheus.String(),
				})
		}
		switch metricDefinition {
		case metricTotalRequests:
			p.totalRequests = metric.(*prometheus.CounterVec)
		case metricRequestsDuration:
			p.requestsDuration = metric.(*prometheus.HistogramVec)
		case metricResponsesSize:
			p.responsesSize = metric.(prometheus.Summary)
		case metricRequestsSize:
			p.requestsSize = metric.(prometheus.Summary)
		}
		metricDefinition.MetricCollector = metric
	}

	p.logger.Logger.Info("Prometheus metrics registered", map[string]interface{}{"from": common.Prometheus.String()})
}

// From https://github.com/DanielHeckrath/gin-prometheus/blob/master/gin_prometheus.go
func getApproxRequestSize(r *http.Request) int {
	s := 0
	if r.URL != nil {
		s = len(r.URL.Path)
	}

	s += len(r.Method)
	s += len(r.Proto)
	for name, values := range r.Header {
		s += len(name)
		for _, value := range values {
			s += len(value)
		}
	}
	s += len(r.Host)

	// r.Form and r.MultipartForm are assumed to be included in r.URL

	if r.ContentLength != -1 {
		s += int(r.ContentLength)
	}
	return s
}

func replaceURLKeys(c *gin.Context) string {
	url := c.Request.URL.Path
	pathUserIDKey := "user_id"

	for _, p := range c.Params {
		if p.Key == pathUserIDKey {
			url = strings.Replace(url, p.Value, ":user_id", 1)
			break
		}
	}
	return url
}
