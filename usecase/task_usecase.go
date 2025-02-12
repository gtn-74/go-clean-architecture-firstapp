package usecase

import (
	"go-clean-architecture-firstapp/model"
	"go-clean-architecture-firstapp/repository"
	"go-clean-architecture-firstapp/validator"
)

type ITaskUsecase interface {
	// taskRepoはmodelから持ってきてる
	//! 第一引数は、引数。model関連は、第二引数としている
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator // !validationを紐付け
}

func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUsecase /* modelのtask.goで定義してる型 */ {
	return &taskUsecase{tr, tv} // taskUsecaseをインスタンス化する際にまとめる
}

func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{} // taskの一覧を格納するためのオブジェクトのスライス

	if err := tu.tr.GetAllTasks(&tasks, userId); /* tasksのメモリアドレスとuserIdを引数として渡してる */ err != nil {
		return nil, err // errの場合、
	}
	resTasks := []model.TaskResponse{} // 返却値として送る中身を変数定義
	// インデックスなくても要素の数で繰り返しできてるから、必要なくていい。
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}

func (tu *taskUsecase) GetTaskById(userId uint, taskId uint) (model.TaskResponse, error) {
	task := model.Task{}
	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

// !validation付与
func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	// !taskRepositoryのcreateTaskを呼び出す前にバリデーションをかける

	// !バリデーション
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}

	// !登録
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {

	// !バリデーション
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}

	// !登録
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	if err := tu.tr.DeleteTask(taskId, userId); err != nil {
		return err
	}
	return nil
}
