package routes

import (
	"net/http"
	"strings"

	"template/infra/httpserver/middlewares"
)

type Method string

const (
	GET    Method = http.MethodGet
	POST   Method = http.MethodPost
	PUT    Method = http.MethodPut
	PATCH  Method = http.MethodPatch
	DELETE Method = http.MethodDelete
)

// Join builds a clean path from parts.
// Join("/api/v2", "samples")        → "/api/v2/samples"
// Join("/api/v2/samples", "{id}")   → "/api/v2/samples/{id}"
// Join("/api/v2/", "/samples")      → "/api/v2/samples"
func Join(parts ...string) string {
	var segments []string

	for _, part := range parts {
		for _, seg := range strings.Split(part, "/") {
			seg = strings.TrimSpace(seg)
			if seg != "" {
				segments = append(segments, seg)
			}
		}
	}

	if len(segments) == 0 {
		return "/"
	}

	return "/" + strings.Join(segments, "/")
}

func Register(
	mux *http.ServeMux,
	method Method,
	base string,
	path string,
	handler http.HandlerFunc,
	mws ...middlewares.Middleware,
) {
	fullPath := Join(base, path)
	mux.Handle(string(method)+" "+fullPath, middlewares.WithMiddleware(handler, mws...))
}
