package data

import (
	"time"

	"github.com/antonisgkamitsios/simple-todo-app/internal/validator"
)

type Todo struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Done      bool
}

type TodoModel struct {
	DB DummyDB
}

type TodoForm struct {
	Title string
	validator.Validator
}

func ValidateTodo(todo *Todo, v *validator.Validator) {
	v.Check(validator.NotBlank(todo.Title), "title", "This field cannot be empty")
	v.Check(validator.MinChars(todo.Title, 10), "title", "This field must be more than 10 characters long")
	v.Check(validator.MaxChars(todo.Title, 150), "title", "This field must be less than 151 characters long")
}

func (m TodoModel) GetAll() ([]Todo, error) {
	return m.DB.Todos, nil
}

func (m TodoModel) Insert(todo *Todo) error {
	id := len(m.DB.Todos) + 1

	todo.ID = int64(id)
	todo.CreatedAt = time.Now()
	todo.Done = false
	m.DB.Todos = append(m.DB.Todos, *todo)

	return nil
}
