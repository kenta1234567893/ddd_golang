package main

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/kenta1234567893/upsider-coding-test/src/cmd/server/libs"
	"github.com/kenta1234567893/upsider-coding-test/src/cmd/server/mymiddleware"
	"github.com/kenta1234567893/upsider-coding-test/src/cmd/server/router"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := setupServer()

	e.Logger.Fatal(e.Start(":1323"))
}

func setupServer() *echo.Echo {
	e := echo.New()

	// ミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = mymiddleware.JSONErrorHandler

	// TODO secret_keyを環境変数から取得するように修正
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret_key"))))

	// ヘルスチェック
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	client := libs.GetClient()

	controllers := libs.InitializeController(client)

	// ルーティング
	router.AuthRouter(e)
	router.ApiRouter(e, controllers)

	return e
}
