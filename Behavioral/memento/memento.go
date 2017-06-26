package memento

// Memento
type memento struct {
	state string
}

func (m *memento) setState(state string) {
	m.state = state
}

func (m *memento) getState() string {
	return m.state
}

// Originator
type originator struct {
	memento *memento
	state string
}

func (o *originator) createMemento() *memento {
	return &memento{o.state}
}

func (o *originator) setState(state string) {
	o.state = state
}

func (o *originator) getState() string {
	return o.state
}

func (o *originator) setMemento(memento *memento) {
	o.memento = memento
}

func (o *originator) restoreToMemento(memento *memento) {
	o.state = memento.getState()
}