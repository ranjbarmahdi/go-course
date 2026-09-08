package httpserver

import (
	"log/slog"
	"net/http"

	"template/infra/httpserver/middlewares"
	"template/infra/runtime"
)

func NewRouter(indicators []runtime.Indicator) http.Handler {
	mux := http.NewServeMux()

	registerHealthRoutes(mux, indicators)

	// route.RegisterRoutes(mux, handler) — per-route middleware inside

	return middlewares.Chain(
		mux,
		middlewares.Recovery,
		middlewares.RequestID,
		middlewares.HttpAccessLog(slog.Default().With("component", "http")),
	)
}
