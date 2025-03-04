package data

type DummyDB struct {
	Todos []Todo
}

type Models struct {
	Todos TodoModel
}

func NewModels(db *DummyDB) Models {
	return Models{
		Todos: TodoModel{DB: db},
	}
}

func NewDummyDB() *DummyDB {
	return &DummyDB{
		Todos: []Todo{{Title: "Clean room"}, {Title: "Change clothes"}},
	}
}
