package prototype

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	product := &concretePrototype{"dumb proto"}
	cloner := product.clone()

	if cloner.GetInfo() != product.GetInfo() {
		t.Errorf("\nProduct:\t[ %s ]\nCloner:\t[ %s ]", product.GetInfo(), cloner.GetInfo())
	}
}