package todo

type Manager struct {
	todos []*ToDo
}

func NewManager() *Manager {

	return &Manager{
		todos: []*ToDo{},
	}
}

func (m *Manager) Add(title string) {
	id := len(m.todos) + 1
	todo := NewToDo(id, title)
	m.todos = append(m.todos, todo)
}

func (m *Manager) List() []*ToDo {
	return m.todos
}

func (m *Manager) Complete(id int) bool {
	for _, todo := range m.todos {
		if todo.ID == id {
			todo.Complete()
			return true
		}
	}
	return false
}
