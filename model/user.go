// user.goがmodelのパッケージに所属してることを明示する必要があるらしい
package model

import "time"

/*
* 下の型をjson形式に変換したタイミングで、""で囲っている文字列をプロパティに置換してくれる
*  `json:"id" gorm:"primaryKey"`
 */

// エンティティ
type User struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Email    string    `json:"email" gome:"unique"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"create_at"` // goのtimeパッケージのtime型を採用
	UpdateAt time.Time `json:"update_at"`
}

// サインアップのエンドポイント
// クライアントにレスポンスで返すデータの型

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
}
