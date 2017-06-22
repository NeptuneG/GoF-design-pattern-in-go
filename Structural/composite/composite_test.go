package composite

import (
	"testing"
)

func TestComposite(t *testing.T) {

	comp1 := &composite{name:"comp1"}
	comp2 := &composite{name:"comp2"}
	comp3 := &composite{name:"comp3"}
	leaf1 := &leaf{"leaf1"}
	leaf2 := &leaf{"leaf2"}

	/*
		comp1
		├─comp2
		│  └─leaf2
		├─leaf1
		└─comp3
	*/
	expect := "comp1-comp2-leaf2-leaf1-comp3"

	comp1.add(comp2)
	comp1.add(leaf1)
	comp2.add(leaf2)
	comp1.add(comp3)
	
	if (expect != comp1.operation()) {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expect, comp1.operation())
	}
}