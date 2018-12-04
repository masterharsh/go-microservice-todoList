package models

import "github.com/jinzhu/gorm"

type (
	// todoModel describes a todoModel type
	TodoModel struct {
		gorm.Model        //this field will embed a Model struct for us which contains four fields “ID, CreatedAt, UpdatedAt, DeletedAt”
		Title      string `json:"title"`
		Completed  int    `json:"completed"`
	}

	// transformedTodo represents a formatted todo
	TransformedTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)
