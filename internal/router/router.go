package router

import "github.com/labstack/echo/v4"

func Init(e *echo.Echo) {
	e.GET("/test", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
}
