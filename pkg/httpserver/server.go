// Package httpserver implements HTTP server.
package httpserver

import (
	"context"
	"net"
	"time"

	"github.com/labstack/echo/v4"
)

// default setting for start http server
const (
	_defaultAddr     = ":80"
	_shutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server          *echo.Echo
	port            string
	notify          chan error
	shutdownTimeout time.Duration
}

// New -.
func New(echo *echo.Echo, port string) *Server {
	s := &Server{
		server:          echo,
		port:            _defaultAddr,
		notify:          make(chan error, 1),
		shutdownTimeout: _shutdownTimeout,
	}
	if port != "" {
		s.port = net.JoinHostPort("", port)
	}
	s.start(port)
	return s
}

func (s *Server) start(port string) {
	go func() {
		s.notify <- s.server.Start(s.port)
		close(s.notify)
	}()

}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}
