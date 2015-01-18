package form

import (
// "github.com/macaron-contrib/binding"
)

type ProjectForm struct {
	Id      int
	Name    string
	Classes []string
}

type TaskForm struct {
	Type      int
	ProjectId int
	Config    string
}
