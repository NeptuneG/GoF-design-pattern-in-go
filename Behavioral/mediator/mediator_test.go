package mediator

import "testing"

func TestMediator(t *testing.T) {
	mediator := &mediator{}
	a := &concretePersonA{mediator}
	b := &concretePersonB{mediator}
	mediator.initRelation(a, b)

	msgA := "msg from a"
	expectAtoB := "msg from a --> person b"
	msgB := "msg from b"
	expectBtoA := "msg from b --> person a"
	
	if a.sendMessage(msgA) != expectAtoB {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expectAtoB, a.sendMessage(msgA))
	}

	if b.sendMessage(msgB) != expectBtoA {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expectBtoA, b.sendMessage(msgB))
	}
}