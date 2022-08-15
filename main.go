package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/kosha/teamwork-connector/docs"
	"github.com/kosha/teamwork-connector/pkg/app"
	"github.com/kosha/teamwork-connector/pkg/logger"
)

var (
	log  = logger.New("app", "teamwork-connector")
	port = 8015
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)

var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "Status of HTTP response",
	},
	[]string{"status"},
)

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_response_time_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)

		statusCode := rw.statusCode

		responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
		totalRequests.WithLabelValues(path).Inc()

		timer.ObserveDuration()
	})
}

func init() {
	prometheus.Register(totalRequests)
	prometheus.Register(responseStatus)
	prometheus.Register(httpDuration)
}

// @title Teamwork Connector API
// @version 1.0
// @description This is a Kosha REST service for exposing many teamwork features as REST APIs with better consistency, observability etc
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email eti@cisco.io
// @host localhost:8015
// @BasePath /
func main() {

	a := app.App{}
	a.Initialize(log)
	a.Router.Use(prometheusMiddleware)

	// Prometheus metrics endpoint
	a.Router.Path("/metrics").Handler(promhttp.Handler())
	log.Infof("Running teamwork-connector on port %d", port)
	a.Run(fmt.Sprintf(":%d", port))

}
