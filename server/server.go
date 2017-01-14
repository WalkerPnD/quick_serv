package server

import (
	"time"

	"github.com/labstack/echo"
)

// Server controles the Running state
type Server struct {
	*echo.Echo
	Running bool
	Port    string
	Root    string
}

// New method creates server with state
func New() *Server {
	server := &Server{
		Running: false,
		Port:    "8080",
		Root:    "./www/",
	}
	server.Echo = echo.New()

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
	if s.Running {
		s.Running = false
		s.Shutdown(1 * time.Second)
	}
}

// ServerHEader adds Header to  "Cache-Control"
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", "max-age=0")
		return next(c)
	}
}
