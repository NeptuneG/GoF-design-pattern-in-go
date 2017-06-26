package strategy

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	
	stg := &strategyA{}
	con := &context{stg}

	expect := "call strategy A interface"
	
	if expect != con.doAction() {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expect, con.doAction())
	}
}