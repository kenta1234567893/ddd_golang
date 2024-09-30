package mymiddleware

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/kenta1234567893/upsider-coding-test/src/shared"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func JSONErrorHandler(err error, c echo.Context) {

	switch e := err.(type) {
	case *shared.MyError:
		e.Print()
		_ = c.JSON(
			e.Status(), e.Error())
	default:
		fmt.Print(err.Error())
		_ = c.JSON(
			http.StatusInternalServerError, err.Error())
	}
}

func CheckSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// セッション取得
		sess, err := session.Get("session", c)
		if err != nil {
			return shared.NewError(http.StatusInternalServerError, "セッションの取得に失敗しました", err)
		}

		// セッションの存在確認
		// TODO 実際はDBなどでセッション情報を確認する
		if sess.Values["session"] != true {
			return shared.NewError(http.StatusUnauthorized, "認証情報の取得に失敗しました。", nil)
		}

		if err = next(c); err != nil {
			return err
		}

		// セッション更新
		// TODO 実際はセッションIDを再生成するなどの対応が必要
		sess.Values["session"] = true
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   60 * 60 * 24, // 1day
			HttpOnly: true,
		}

		if err = sess.Save(c.Request(), c.Response()); err != nil {
			return shared.NewError(http.StatusInternalServerError, "セッションの設定に失敗しました。", err)
		}

		return nil
	}
}
