package state

// State
type state interface {
	opeChangeState(*context)
	callInterface(*context) string
	changeState(con *context, st state) bool
}

type concreteStateA struct {
}

func (ca *concreteStateA) opeChangeState(context *context) {
	ca.changeState(context, &concreteStateB{})
}

func (ca *concreteStateA) callInterface(context *context) string {
	return "concreteStateA.callInterface()"
}

func (ca *concreteStateA) changeState(context *context, state state) bool {
	return context.changeState(state)
}

type concreteStateB struct {
}

func (cb *concreteStateB) opeChangeState(context *context) {
	cb.changeState(context, &concreteStateA{})
}

func (cb *concreteStateB) callInterface(context *context) string {
	return "concreteStateB.callInterface()"
}

func (cb *concreteStateB) changeState(context *context, state state) bool {
	return context.changeState(state)
}

type context struct {
	state state
}

func (c *context) opeChangeState() {
	c.state.opeChangeState(c)
}

func (c *context) callInterface() string {
	return c.state.callInterface(c)
}

func (c *context) changeState(state state) bool {
	c.state = state
	return true
}