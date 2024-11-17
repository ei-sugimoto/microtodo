package cmd

import (
	"log/slog"

	"github.com/ei-sugimoto/microtodo/server/bff/client"
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
	Server.Routing()
	Server.e.Use(middleware.Logger())
	Server.e.Use(middleware.Recover())
	slog.Info("Server started")
	Server.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	Server.e.Start(":5555")
}

func (s *Server) Routing() {
	s.e.GET("/health", client.NewHealthClient().Health)
	s.e.POST("/member/create", client.NewMemberClient().Create)
	s.e.POST("/member/login", client.NewMemberClient().Login)
}
