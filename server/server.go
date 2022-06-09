package server

import "github.com/labstack/echo/v4"

type Server interface {
	AddGroup(path string) *echo.Group
	Serve(port string)
}
