package repository

import (
	"fmt"
	"go-clean-architecture-firstapp/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// taskとの紐付け
type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error           // taskの一覧を配列に格納するために第一引数に、model.taskのスライスのポインターを格納
	GetTaskById(task *model.Task, userId uint, taskId uint) error // taskIdに一致するタスクを取得する
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, taskId uint, userId uint) error
	DeleteTask(userId uint, taskId uint) error
}

// dbフィールドを持ったタスク関連のDB操作を行う、オブジェクトのこと
type taskRepository struct {
	db *gorm.DB
}

// 外側でインスタンス化されたdbを引数として受け取ってる
func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db} // 受け取ったdbを使って、構造体の実態を作ってる
}

// getAllTasks
// taskレシーバーをポインタレシーバーとして受け取る形でgetAllTasksとして定義してる
// レシーバーとは：構造体のポインタのこと。
func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error { // 引数の型は、インターフェースと一緒にする必要がある
	// taskRepoをtrとしてる
	// dbの中でuserIdが引数のuserIdと一致するtasksの一覧を取得してる。orderで新しく作成されたものを一番下に表示してる。
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil // エラーが無いを返してるだけ。
}

func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	// 引数と一致するtaskIdをtaskの中から取得する
	if err := tr.db.Joins("User").Where("user_id", userId).First(task, taskId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	// taskオブジェクトのポインタを渡す
	// taskModelから引数で受け取った、taskIdとuserIdと一途するtitleを変更してくださいという処理
	//! Clauses(clause.Returning{})まじでよくわからん
	// ! Clauses(clause.Returning{}) は、更新された行のデータ取得
	// Returning{}のデフォ挙動は、全てのカラムの値を返す
	//! updateで使うと、更新後のデータを即取得できる
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskId, userId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	// resultと関連するレコードが1未満だった時のエラー
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	//! この書き方は、決まってる Delete(&構造体名{}) Delete(&model.Task{})
	result := tr.db.Where("id=? AND user_id=?", userId, taskId).Delete(&model.Task{})
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
