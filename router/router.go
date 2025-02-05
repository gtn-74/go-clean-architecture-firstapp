package router

import (
	"go-clean-architecture-firstapp/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

// taskControllerを使用できるように実装
func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	// userControllerを引数で受け取れるようにする
	// echoのインスタンスを作成する
	// エンドポイント追加
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.LogOut)

	t := e.Group("/tasks")
	// middleware echoのjwtというミドルウェア
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token", // cookieの中にトークンという名前で保存
	}))

	// task関係のエンドポイントにリクエストがあった場合のコントローラーを呼び出すようにしてる
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)
	return e
}
