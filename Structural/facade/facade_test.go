package facade

import (
	"testing"
)

func TestFacade(t *testing.T) {

	sys1 := &subsys1{}
	sys2 := &subsys2{}
	facade := &facade{sys1, sys2}
	expect := "wrapper: [foobar]"

	if expect != facade.wrapper() {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expect, facade.wrapper())
	}
}