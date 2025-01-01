package main

import (
	"fmt"
	"go-clean-architecture-firstapp/db"
	"go-clean-architecture-firstapp/model"
)

func main() {
	dbConn := db.NewDB() // dbConnは、最初作ったDB
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{}) // modelはテーブル情報を書いたファイル
}
