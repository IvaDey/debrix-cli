package dbUtils

import (
	"gorm.io/gorm"
)

type StoredTodo struct {
	gorm.Model
	TodoItem
}

type TodoItem struct {
	RelativePath string
	FileName     string
	Line         uint32
	Task         string
	Scope        string
	Assignee     string
	Due          string
	Reminder     string

	IsCompleted bool
	Author      string
}
