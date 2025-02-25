package data

import "time"

type Todo struct {
	ID          int64
	CreatedAt   time.Time
	Title       string
	Description string
}

type TodoModel struct {
	DB DummyDB
}

func (m TodoModel) GetAll() ([]Todo, error) {
	return m.DB.Todos, nil
}
