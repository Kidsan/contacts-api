package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (h *HTTPServer) buildMetricsRouter() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", promhttp.Handler().ServeHTTP)
	}
}
