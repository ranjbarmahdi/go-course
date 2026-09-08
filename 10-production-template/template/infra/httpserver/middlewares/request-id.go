package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const RequestIDHeader = "x-request-id"

type requestIDContextKey struct{}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(RequestIDHeader)
		if id == "" {
			id = uuid.NewString()
		}

		w.Header().Set(RequestIDHeader, id)

		ctx := context.WithValue(r.Context(), requestIDContextKey{}, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequestIDFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(requestIDContextKey{}).(string)
	return id, ok
}

func requestIDOf(ctx context.Context) string {
	id, _ := RequestIDFromContext(ctx)
	return id
}
