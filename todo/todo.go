package todo

import "time"

type ToDo struct {
	ID        int
	Title     string
	Completed bool
	Created   time.Time
}

func NewToDo(id int, title string) *ToDo {
	return &ToDo{
		ID:        id,
		Title:     title,
		Completed: false,
		Created:   time.Now(),
	}
}

func (t *ToDo) Complete() {
	t.Completed = true
}
