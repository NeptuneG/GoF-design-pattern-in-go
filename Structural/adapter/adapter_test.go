package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {

	expect := "adaptee's request & response"

	adaptee := &adaptee{}
	adapter := &adapter{adaptee}

	if callTarget(adapter) != expect {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expect, callTarget(adapter))
	}
}

func callTarget(target target) string {
	return target.request()
}
