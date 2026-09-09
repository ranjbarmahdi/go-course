package httpserver

import (
	"log/slog"
	"net/http"

	"template/infra/httpserver/middlewares"
	samplehttp "template/infra/httpserver/routes/sample"
	"template/infra/runtime"
)

const APIV2 = "/api/v2"
const ManagementV2 = "/api/v2/management"

func NewRouter(indicators []runtime.Indicator, sampleHandler *samplehttp.Handler) http.Handler {
	mux := http.NewServeMux()

	registerHealthRoutes(mux, indicators)

	samplehttp.RegisterRoutes(mux, APIV2, sampleHandler)

	return middlewares.Chain(
		mux,
		middlewares.RequestID,
		middlewares.Recovery,
		middlewares.HttpAccessLog(slog.Default().With("component", "http")),
		middlewares.StripTrailingSlash,
	)
}
