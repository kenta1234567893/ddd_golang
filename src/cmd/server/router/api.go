package router

import (
	"github.com/kenta1234567893/upsider-coding-test/src/cmd/server/mymiddleware"
	"github.com/labstack/echo/v4"
)

func ApiRouter(e *echo.Echo) {

	authed := e.Group("")
	authed.Use(mymiddleware.CheckSession)

}
