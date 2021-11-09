package model

import "time"

type Todo struct {
	Item        string    `json:"item" bson:"item"`
	Time        time.Time `json:"time" bson:"time"`
	LastUpdated time.Time `json:"lastUpdated,omitempty" bson:"lastUpdated,omitempty"`
}

func (todo *Todo) IsEmpty() bool {
	return todo.Item == ""
}
