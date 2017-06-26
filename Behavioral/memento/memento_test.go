package memento

import (
	"testing"
)

func TestMemento(t *testing.T) {
	o := &originator{}
	o.setState("old")
	if "old" != o.getState() {
		t.Errorf("originator init state error")
	}

	m := o.createMemento()
	o.setState("new")
	if "new" != o.getState() {
		t.Errorf("originator set state error")
	}	

	o.restoreToMemento(m)
	if "old" != o.getState() {
		t.Errorf("originator restore state error")
	}
}