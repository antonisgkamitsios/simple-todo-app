package data

import "time"

type Todo struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Done      bool
}

type TodoModel struct {
	DB DummyDB
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
