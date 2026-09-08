package middlewares

import (
	"log/slog"
	"net"
	"net/http"
	"strings"
	"time"
)

type statusCapturingWriter struct {
	http.ResponseWriter

	status      int
	bytes       int
	wroteHeader bool
}

func (w *statusCapturingWriter) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}

func (w *statusCapturingWriter) WriteHeader(status int) {
	if w.wroteHeader {
		return
	}

	w.status = status
	w.wroteHeader = true

	w.ResponseWriter.WriteHeader(status)
}

func (w *statusCapturingWriter) Write(data []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}

	n, err := w.ResponseWriter.Write(data)
	w.bytes += n

	return n, err
}

func HttpAccessLog(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !logger.Enabled(r.Context(), slog.LevelDebug) {
				next.ServeHTTP(w, r)
				return
			}

			start := time.Now()

			writer := &statusCapturingWriter{
				ResponseWriter: w,
				status:         http.StatusOK,
			}

			next.ServeHTTP(writer, r)

			logger.DebugContext(
				r.Context(),
				"http request",
				slog.String("request_id", requestIDOf(r.Context())),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status", writer.status),
				slog.Int("bytes", writer.bytes),
				slog.Duration("duration", time.Since(start)),
				slog.String("client_ip", remoteIP(r)),
				slog.String("user_agent", r.UserAgent()),
				slog.String("referer", r.Referer()),
			)
		})
	}
}

func remoteIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err == nil {
		return host
	}

	return strings.TrimSpace(r.RemoteAddr)
}
