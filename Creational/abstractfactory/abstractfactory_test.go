package abstractfactory

import "testing"

func TestAbstractFactory(t *testing.T) {

	factory := &concreateFactory{}

	prod1 := factory.createProductA()
	prod2 := factory.createProductB()

	if prod1.info() != "Product A" || prod2.info() != "Product B" {
		t.Errorf("\nExpect:\t[ Product A ][ Product B ]\nActual:\t[ %s ][ %s ]", prod1.info(), prod2.info())
	}
}
