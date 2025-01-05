package router

import (
	"go-clean-architecture-firstapp/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	// userControllerを引数で受け取れるようにする
	// echoのインスタンスを作成する
	// エンドポイント追加
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.LogOut)
	return e
}
