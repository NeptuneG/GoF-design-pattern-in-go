package observer

import (
	"fmt"
)

// observer
type observer interface {
	update(subject)
	getSubject() subject
}

type concreteOberver struct {
	subject subject
	state string
}

func (co *concreteOberver) getSubject() subject {
	return co.subject
}

func (co *concreteOberver) update(subject subject) {
	co.state = subject.getState()
	co.notifyResponser()
}

func (co *concreteOberver) notifyResponser() {
	fmt.Println("got notified: " + co.subject.getState())
}

