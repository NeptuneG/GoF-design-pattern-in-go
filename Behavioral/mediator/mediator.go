package mediator

// Mediator
type mediator struct {
	a person
	b person
}

func (m *mediator) send(msg string, person person) string {
	if (person == m.a) {
		return msg + " --> person b"
	} else if (person == m.b) {
		return msg + " --> person a"
	}
	return ""
}

func (m *mediator) initRelation(a person, b person) {
	m.a = a
	m.b = b
}

// Person
type person interface {
	sendMessage(msg string) string
}

type concretePersonA struct {
	m *mediator
}

func (ca *concretePersonA) sendMessage(msg string) string {
	return ca.m.send(msg, ca)
}

type concretePersonB struct {
	m *mediator
}

func (cb *concretePersonB) sendMessage(msg string) string {
	return cb.m.send(msg, cb)
}
