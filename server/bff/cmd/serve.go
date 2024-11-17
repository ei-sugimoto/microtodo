package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	e *echo.Echo
}

func Serve() {
	Server := &Server{
		e: echo.New(),
	}

	Server.e.Use(middleware.Logger())
	Server.e.Use(middleware.Recover())
	Server.e.Start(":5555")
}

func (s *Server) Routing() {

}
