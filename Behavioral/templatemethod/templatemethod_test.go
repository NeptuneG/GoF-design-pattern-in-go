package templatemethod

import (
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	ope1 := &concreteOpe1{}
	ope2 := &concreteOpe2{}

	expect1 := "Ope1.Ope1()|Ope1.Ope2()"
	expect2 := "Ope2.Ope1()|Ope2.Ope2()"

	if ope1.templateMethod(ope1) != expect1 {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expect1, ope1.templateMethod(ope1))
	}

	if ope2.templateMethod(ope2) != expect2 {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expect2, ope2.templateMethod(ope2))
	}
}