package httpserver

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"time"

	"template/infra/runtime"
)

type Addr string

const (
	shutdownTimeout   = 5 * time.Second
	readHeaderTimeout = 5 * time.Second
	readTimeout       = 15 * time.Second
	writeTimeout      = 15 * time.Second
	idleTimeout       = 60 * time.Second
)

type Server struct {
	srv *http.Server
}

var log = slog.Default().With("component", "http")

func NewServer(addr Addr, handler http.Handler, ctx context.Context) *Server {
	return &Server{
		srv: &http.Server{
			Addr:              string(addr),
			Handler:           handler,
			ReadHeaderTimeout: readHeaderTimeout,
			ReadTimeout:       readTimeout,
			WriteTimeout:      writeTimeout,
			IdleTimeout:       idleTimeout,
			BaseContext: func(_ net.Listener) context.Context {
				return ctx
			},
		},
	}
}

func (s *Server) Name() string { return "http" }

func (s *Server) Start(ctx context.Context) error {
	if ctx.Err() != nil {
		return nil
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- s.srv.ListenAndServe()
	}()

	log.Info("listening", "addr", s.srv.Addr)

	select {
	case err := <-errCh:
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	case <-ctx.Done():
		log.Info("shutdown started")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		shutdownErr := s.Stop(shutdownCtx)
		serveErr := <-errCh

		if shutdownErr != nil {
			return shutdownErr
		}
		if !errors.Is(serveErr, http.ErrServerClosed) {
			return serveErr
		}
		log.Info("shutdown successfully")
		return nil
	}
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

var _ runtime.Runner = (*Server)(nil)
