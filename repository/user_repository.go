package repository

import (
	"go-clean-architecture-firstapp/model"

	"gorm.io/gorm"
)

// インターフェースを定義する
// インターフェースを可視化するために、Iを先頭につける
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
	// *は、ポインタ
}

// goのインターフェースは、メソッドの集まりを定義する
// メソッドを定義するためのもの

// 具体的なrepositoryのソースコードを書く
type userRepository struct {
	db *gorm.DB
}

// dbから取得した値とIUserRepositoryで利用できるようにする
// そのために、GetUserByEmailと、CreateUserのメソッドを実装する必要がある。
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// 引数の型と、返り値の型は、インターフェースのメソッドと同じにする必要がある。
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	// dbの中で、emailの値が、引数で受け取ったemailと同じuserを探す
	// emailが一致したら、userで受け取っている内容を利用する
	// 失敗したらエラー、成功したらnilを返却
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	// 入力値をuserとして新たに登録作成する。
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
