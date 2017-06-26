package state

import (
	"testing"
)

func TestState(t *testing.T) {
	
	state := &concreteStateA{}
	context := &context{state}

	expectA := "concreteStateA.callInterface()"
	expectB := "concreteStateB.callInterface()"

	if expectA != context.callInterface() {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expectA, context.callInterface())
	}
	context.opeChangeState()
	if expectB != context.callInterface() {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expectB, context.callInterface())
	}
}