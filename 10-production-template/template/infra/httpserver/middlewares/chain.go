package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

// Chain(h, a, b) => a(b(h)) — first listed is outermost.
func Chain(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func WithMiddleware(handler http.HandlerFunc, middlewares ...Middleware) http.Handler {
	var h http.Handler = handler
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
