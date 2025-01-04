package repository

import "go-clean-architecture-firstapp/model"

// インターフェースを定義する
// インターフェースを可視化するために、Iを先頭につける
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
	// *は、ポインタ
}

// goのインターフェースは、メソッドの集まりを定義する
// メソッドを定義するためのもの
