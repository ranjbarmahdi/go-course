package httpserver

import (
	"net/http"

	"template/infra/runtime"
)

func registerHealthRoutes(mux *http.ServeMux, indicators []runtime.Indicator) {
	mux.HandleFunc("GET /livez", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		for _, indicator := range indicators {
			if !indicator.Ready(r.Context()) {
				http.Error(w, indicator.Name()+" not ready", http.StatusServiceUnavailable)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
	})
}
