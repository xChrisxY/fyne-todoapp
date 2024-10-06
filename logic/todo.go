package models

import (
	"fmt"
)

type Todo struct {
	Descripcion string 
	Done bool
}

func NewTodo(description string) Todo {
	return Todo{Descripcion: description, Done: false}
}

func (t Todo) String() string {
	fmt.Println(t.Descripcion, t.Done)
	return fmt.Sprintf(t.Descripcion, t.Done)
}

