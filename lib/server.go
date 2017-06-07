package lib

import (
	"context"
	"time"

	"github.com/labstack/echo"
)

// Server controles the Running state
type Server struct {
	*echo.Echo
	Running bool
	*Config
}

// NewServer method creates server with state
func NewServer() *Server {
	server := &Server{
		Running: false,
	}
	server.Echo = echo.New()
	server.Config = LoadConfig()

	return server
}

// Run server if it's not running
func (s *Server) Run() {
	if !s.Running {
		s.Running = true
		s.Static("/", s.Root)

		s.Use(serverHeader)
		go func() {
			s.Logger.Debug(s.Start(":" + s.Port))
		}()
	}
}

// Stop server if it's running
func (s *Server) Stop() {
	if !s.Running {
		return
	}
	s.Running = false
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		s.Logger.Fatal(err)
	}
}

// ServerHEader adds Header to  "Cache-Control"
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", "max-age=0")
		return next(c)
	}
}
