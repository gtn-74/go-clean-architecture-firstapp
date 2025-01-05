# クリーンアーキテクチャで REST API

## クリーンアーキテクチャ(OOP:Object-Oriented Programming)

オブジェクト指向の概念は取り入れているものの、オブジェクト指向に限定されるアーキテクチャではない。

システムを以下のような層に分けて、依存関係を明確にし、変更に強いシステムを設計する方法

- 中心（エンティティ層）
  ビジネスルールを明確にし、アプリケーションに依存しない

- ユースケース層
  アプリケーション固有のビジネスロジックを定義

- インターフェースアダプタ層
  外部システム（データベース、API）ブリッジ

- フレームワーク層
  フレームワーク、外部ライブラリを使った具体的な実装

### クリーンアーキテクチャの基本原則

依存逆転の原則（DIP:Dependency Inversion Principle）を実現することが目的

#### オブジェクト指向の特徴

- カプセル化
- 継承
- ポリモーフィズム：同じインターフェースを持つ異なる実装を切り替える仕組み

## go のエトセトラ

go mod = Go Modules の省略であり、依存関係管理システムのこと

```go
go mod init
```

### go.mod の内容

- モジュール名
- Golang のバージョン
- 依存関係

### 依存関係の整理

`go mod tidy`

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

pgAdmin で生成したテーブルを確認できる。

![スクリーンショット 2025-01-01 11 20 58](https://gist.github.com/user-attachments/assets/1fa320ab-d526-49be-9c32-1952fd68e648)

## golang の独特記法

### 短縮変数宣言

下の`:=`と`var`の変数宣言はイコール

```go
:= === var hoge string = fuga
```

### method receiver

```go
func (レシーバー名 レシーバー型) メソッド名(引数) 戻り値型 {
    // メソッドの処理
}

func (uu *userUsecase)
```

### CRUD チェック

- docker-compose 立ち上げ
- go run migrate 'GO_ENV=dev go run migrate/migrate.go'
- echo run 'GO_ENV=dev go run main.go'

## あとで調べよう

### ポインタ

### レシーバー
