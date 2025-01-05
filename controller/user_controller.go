package controller

import (
	"go-clean-architecture-firstapp/model"
	"go-clean-architecture-firstapp/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

// 引数として、echoで受け取れるContextを定義
type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	LogOut(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	// userControllerが、IUserControllerを満たすためには、３つのメソッドを定義する必要がある
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	// クライアントから受け取るリクエストbodyの値を構造体に変換する処理
	// 空の構造体を定義
	user := model.User{}
	// echoのContextに存在するBindメソッドを使う
	if err := c.Bind(&user); err != nil {
		// Bind失敗処理
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// Bind成功以降
	userRes, err := uc.uu.SignUp(user)
	// signUp失敗
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error()) // StatusInternalServerErrorは、500
	}
	// signUp成功
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) Login(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user) //userUsecaseのloginはトークンを発行する
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// !ではない。
	// return c.JSON(http.StatusCreated, tokenString)
	// cookieにトークンを持たせる処理が必要

	cookie := new(http.Cookie) //httpに定義されているcookie構造体を変数化
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// cookie.Secure = true ポストマンで確認するまでは、false設定
	cookie.HttpOnly = true                  // クライアントjsからトークンの値を読み取れない設定
	cookie.SameSite = http.SameSiteNoneMode //front,backのドメインが異なる。クロスドメイン間のcookie送受信定義
	c.SetCookie(cookie)                     // httpレスポンスに、cookieの内容を含める
	return c.NoContent(http.StatusOK)
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// cookie.Secure = true ポストマンで確認するまでは、false設定
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
