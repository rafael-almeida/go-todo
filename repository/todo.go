package repository

import (
	"fmt"

	"github.com/rafael-almeida/go-todo/model"
)

type todo struct {
	Id        string
	CreatedAt string
}

var todos = []todo{{Id: "1", CreatedAt: "1/1/2021"}, {Id: "2", CreatedAt: "1/2/2021"}}

func Scan() []model.Todo {
	var ts []model.Todo

	for _, t := range todos {
		todo := model.Todo{Id: t.Id}
		ts = append(ts, todo)
	}

	return ts
}

func Get(id string) *model.Todo {
	for _, t := range todos {
		if t.Id == id {
			return &model.Todo{Id: t.Id}
		}
	}

	return nil
}

func Put(t *model.Todo) bool {
	prevLen := len(todos)

	td := todo{Id: fmt.Sprint(prevLen + 1), CreatedAt: "1/2/2021"}
	todos = append(todos, td)

	if len(todos) != (prevLen + 1) {
		return false
	}

	return true
}

func Delete(id string) *model.Todo {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return &model.Todo{Id: t.Id}
		}
	}

	return nil
}
