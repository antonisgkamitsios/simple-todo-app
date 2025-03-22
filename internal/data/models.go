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
		Todos: []Todo{{ID: 1, Title: "Clean room"}, {ID: 2, Title: "Change clothes"}},
	}
}
