package model

type TaskType string

const (
	TaskTypeCreate TaskType = "task_type_create"
	TaskTypeUpdate TaskType = "task_type_update"
)

type Task struct {
	Type TaskType
	Data interface{}
}

func NewTask(taskType TaskType, data interface{}) Task {
	return Task{
		Type: taskType,
		Data: data,
	}
}
