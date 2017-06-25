package command

import (
	"testing"
)

func TestCommand(t *testing.T) {

	invoker := &invoker{&concreteCommand{&reciever{"foobar"}}}
	expect := "reciever.action(): foobar"

	if invoker.invoke() != expect {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expect, invoker.invoke())
	}
}
