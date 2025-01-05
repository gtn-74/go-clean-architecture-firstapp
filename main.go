package main

import (
	"go-clean-architecture-firstapp/controller"
	"go-clean-architecture-firstapp/db"
	"go-clean-architecture-firstapp/repository"
	"go-clean-architecture-firstapp/router"
	"go-clean-architecture-firstapp/usecase"
)

func main() {
	// dbインスタンス化
	db := db.NewDB()
	// repositoryで作成したコンストラクタを起動
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController((userUsecase))
	e := router.NewRouter((userController))
	e.Logger.Fatal(e.Start(":8080"))

}
