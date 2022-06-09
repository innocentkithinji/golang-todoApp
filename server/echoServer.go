package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"log"
)

type server struct {
	echo *echo.Echo
}

func (s server) AddGroup(path string) *echo.Group {
	group := s.echo.Group(path)

	return group
}

func (s server) Serve(port string) {
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.Logger())
	serviceName := viper.Get("service_name").(string)
	env := viper.Get("environment").(string)
	s.echo.Logger.SetPrefix(serviceName + ":" + env)

	log.Fatal(s.echo.Start(":" + port))
}

func NewServer() Server {
	e := echo.New()
	return &server{echo: e}
}
