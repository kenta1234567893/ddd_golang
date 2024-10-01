package router

import (
	"github.com/kenta1234567893/upsider-coding-test/src/cmd/server/libs"
	"github.com/kenta1234567893/upsider-coding-test/src/cmd/server/mymiddleware"
	"github.com/labstack/echo/v4"
)

func ApiRouter(e *echo.Echo, controllers *libs.InitializeControllers) {

	authed := e.Group("")
	authed.Use(mymiddleware.CheckSession)

	authed.POST("/api/invoices", controllers.InvoiceController.Issue)
	authed.GET("/api/invoices", controllers.InvoiceController.Search)

}
