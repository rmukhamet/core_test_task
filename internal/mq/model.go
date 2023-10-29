package mq

import (
	"encoding/json"
	"time"

	"github.com/rmukhamet/core_test_task/internal/model"
)

type Task struct {
	Type string          `json:"task_type"`
	Data json.RawMessage `json:"data"`
}

type Retailer struct {
	// constant fields
	ID      string
	Name    string
	Address Address
	// keep version
	Owner     Person
	OpenTime  time.Time
	CloseTime time.Time
	Version   Version
}

// func (r Retailer) Validate() error {
// 	// TODO
// 	return nil
// }

type Address struct {
	City   string
	Street string
	House  string
}

type Person struct {
	FirstName string
	LastName  string
}

type Version struct {
	Actor   string
	Version int
}

func NewTask(t model.Task) Task {
	value, _ := json.Marshal(t.Data)
	return Task{
		Type: string(t.Type),
		Data: value,
	}
}
func (t Task) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(t)
	return bytes, err
}

func (t Task) ToDTO() model.Task {
	return model.Task{
		Type: model.TaskType(t.Type),
		Data: t.Data,
	}
}
