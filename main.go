package main

import (
	"go-clean-architecture-firstapp/controller"
	"go-clean-architecture-firstapp/db"
	"go-clean-architecture-firstapp/repository"
	"go-clean-architecture-firstapp/router"
	"go-clean-architecture-firstapp/usecase"
	"go-clean-architecture-firstapp/validator"
)

func main() {
	// dbインスタンス化
	db := db.NewDB()

	// validation
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()

	// repositoryで作成したコンストラクタを起動
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)

	taskController := controller.NewTaskController((taskUsecase))
	userController := controller.NewUserController((userUsecase))

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))

}
