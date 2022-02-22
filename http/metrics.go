package http

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (h *HTTPServer) handleMetrics() http.HandlerFunc {
	return promhttp.Handler().ServeHTTP
}
