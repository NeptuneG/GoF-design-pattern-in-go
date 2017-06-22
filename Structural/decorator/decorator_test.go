package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	
	comp := &concreteComponent{}
	dec := &decorator{comp}

	if comp.operation() != "dumb comp" || dec.operation() != "###dumb comp###" {
		t.Errorf("\nExpect:\t[ dumb comp ][ ###dumb comp### ]\nActual:\t[ %s ][ %s ]", comp.operation(), dec.operation())
	}
	
}