package db

import (
	"fmt"
	"log"
	"os"

	// godotenvは、外部パッケージのため、`go mod tidy`コマンドでインストールしておく必要がある
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// ここにDBの初期化処理を書く
	// 環境変数を読み込むための処理
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load() // .envファイルを読み込む
		if err != nil {
			log.Fatal(err)
		}
	}
	// POSTGRES_USER, POSTGRES_PW, POSTGRES_HOST, POSTGRES_PORT, POSTGRES_DBは、.envファイルに書いてる。
	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	// urlとdbを接続
	db, err := gorm.Open((postgres.Open(url)), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(("Connceted")) // 接続
	return db
}
// gormは、外部パッケージのため、`go mod tidy`コマンドでインストールしておく必要がある

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatal(err)
	}
}