package factorymethod

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	
	expect := "a concrete product"
	
	factory := &concreateFactory{}

	if (factory.createProduct().info() != expect) {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expect, factory.createProduct().info())
	}
	
}