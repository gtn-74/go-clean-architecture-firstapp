# クリーンアーキテクチャで REST API

## go mod init

```go
go mod init
```

## docker command

`web`、`app`2 つのコンテナを起動:`docker up -d`

※コンテナは裏側で動いてるため、ターミナルでは動いて見えない
※`-d`がなくても dockerDesktop を開いていれば、立ち上がるが、ターミナルを一つ使うことになる。

コンテナ一覧:`docker ps`

```go
type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json: "email" gorm:"unique"`
                     // 文字列がズレる。インデントチェックした方が良いかも
}
```

![スクリーンショット 2025-01-01 11 20 58](https://gist.github.com/user-attachments/assets/b5875fbc-8e1b-4208-908b-e72c2d4f3b05)


`GO_ENV`の値が`dev`だったら load 関数が走り、ローカルで持ってる`.env`を見に行ってる。

```go
func NewDB() *gorm.DB {
	// ここにDBの初期化処理を書く
	// 環境変数を読み込むための処理
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load() // .envファイルを読み込む
		if err != nil {
			log.Fatal(err)
		}
	}
}
```

### 注意
docker も立ち上げてないと動かないからね。

`go run migrate/migrate.go`したら、

pgAdminで生成したテーブルを確認できる。

![スクリーンショット 2025-01-01 11 20 58](https://gist.github.com/user-attachments/assets/1fa320ab-d526-49be-9c32-1952fd68e648)