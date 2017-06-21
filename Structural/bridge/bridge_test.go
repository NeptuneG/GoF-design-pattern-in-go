package bridge

import (
	"testing"
)

func TestBridge(t *testing.T) {
	expect := "concrete abstaction implement"
	concrete := &concreteAbstractionImpl{}
	refined := &refinedAbstraction{concrete}

	if expect != callOperation(refined) {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expect, callOperation(refined))
	}
}

func callOperation(a abstraction) string {
	return a.operation()
}