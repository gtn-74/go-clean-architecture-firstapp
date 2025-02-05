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
	taskRepository := repository.NewTaskRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository)
	taskUsecase := usecase.NewTaskUsecase((taskRepository))

	taskController := controller.NewTaskController((taskUsecase))
	userController := controller.NewUserController((userUsecase))

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))

}
