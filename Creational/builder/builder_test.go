package builder

import "testing"

func TestBuilder(t *testing.T) {

	expect := "CPU:3.30GHz-RAM:8GB-Disk:500GB"

	builder := &dumbBuilder{}
	director := director{builder}
	model := director.Constructor()
	product := builder.getComputer(model)

	if product.model != expect {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expect, product.model)
	}
}
