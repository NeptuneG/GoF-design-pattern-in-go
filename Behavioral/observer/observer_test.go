package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	
	sub := &concreteSubject{}
	o1 := &concreteOberver{subject:sub}
	o2 := &concreteOberver{subject:sub}

	sub.attach(o1)
	sub.attach(o2)

	sub.setState("set state to foo")
	sub.notify()
	sub.setState("set state to bar")
	sub.notify()

}