package visitor

import (
	"testing"
)

func TestVisitor(t *testing.T) {
	visitor := &concreteVisitorA{}
	element := &concreteElementB{}
	
	expect := "eleB is visited"

	if element.accept(visitor) != expect {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expect, element.accept(visitor))
	}
}