package data

import (
	"errors"
	"slices"
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
	DB *DummyDB
}

type TodoForm struct {
	Title               string `form:"title"`
	validator.Validator `form:"-"`
}

func ValidateTodo(todo *Todo, v *validator.Validator) {
	v.Check(validator.NotBlank(todo.Title), "title", "This field cannot be empty")
	v.Check(validator.MinChars(todo.Title, 10), "title", "This field must be more than 10 characters long")
	v.Check(validator.MaxChars(todo.Title, 150), "title", "This field must be less than 150 characters long")
}

func (m TodoModel) GetAll() ([]Todo, error) {
	cpyTodos := slices.Clone(m.DB.Todos)

	slices.SortFunc(cpyTodos, func(a, b Todo) int {

		return b.CreatedAt.Compare(a.CreatedAt)
	})
	return cpyTodos, nil
}

func (m TodoModel) Insert(todo *Todo) error {
	id := len(m.DB.Todos) + 1

	todo.ID = int64(id)
	todo.CreatedAt = time.Now()
	todo.Done = false
	m.DB.Todos = append(m.DB.Todos, *todo)

	return nil
}

func (m TodoModel) Delete(id int64) error {
	todoIndexToDelete := -1

	for index, todo := range m.DB.Todos {
		if id == todo.ID {
			todoIndexToDelete = index
			break
		}
	}
	if todoIndexToDelete == -1 {
		return errors.New("could not find todo to delete")
	}

	m.DB.Todos = slices.Delete(m.DB.Todos, todoIndexToDelete, todoIndexToDelete+1)
	return nil
}
