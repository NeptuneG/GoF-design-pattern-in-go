package chainofresposibility

import "testing"

func TestChainOfResponsibility(t *testing.T) {
	a := &concreteHandlerA{&baseHandler{info: "handlerA"}}
	b := &concreteHandlerA{&baseHandler{info: "handlerB"}}
	c := &concreteHandlerB{&baseHandler{info: "handlerB"}, "foorbar"}

	a.setSuccessor(b)
	b.setSuccessor(c)

	request := "zepp"
	expect := request + "->[BaseHandler]Default->[BaseHandler]Default->[ConcHandlerB]foorbar"

	actual := a.handleRequest(request)
	if expect != actual {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expect, actual)
	}
}
