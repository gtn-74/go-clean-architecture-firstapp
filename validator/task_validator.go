package validator

import (
	"go-clean-architecture-firstapp/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// インターフェース
// バリデーター
// 構造体を定義する

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct{} //  taskValidatorがITaskValidatorを満たすための中身を記述

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

// taskValidatorをレシーバーとして定義
func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(&task.Title,
			validation.Required.Error("title is required"),            // 必須条件のバリデーション
			validation.RuneLength(1, 10).Error("limited max 10 char"), // 文字数制限バリデーション
		))
}
