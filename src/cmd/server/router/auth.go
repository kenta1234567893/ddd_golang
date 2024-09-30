package router

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/kenta1234567893/upsider-coding-test/src/shared"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AuthRouter(e *echo.Echo) {
	e.POST("/login", func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return shared.NewError(http.StatusInternalServerError, "セッションの取得に失敗しました。", err)
		}

		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   60 * 60 * 24, // 1day
			HttpOnly: true,
		}

		sess.Values["session"] = true
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return err
		}
		return c.String(http.StatusOK, "success")
	})

	e.POST("/logout", func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return shared.NewError(http.StatusInternalServerError, "セッションの取得に失敗しました。", err)
		}
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		}
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return err
		}
		return c.String(http.StatusOK, "success")
	})

}
