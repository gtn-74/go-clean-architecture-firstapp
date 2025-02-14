package model

import "time"

type Task struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Title    string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	User     User      `json:"user" gorm:"foreignKey:UserId; constraint: OnDelete:CASCADE"`
	UserId   uint      `json:"user_id" gorm:"not null"`
}

// Userのprimary keyがUserIdと紐づくことで1対多の関係にできる ID       uint      `json:"id" gorm:"primaryKey"`

// constraint: OnDelete:CASCADE でUserが削除されたらTaskも削除される

type TaskResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
}
