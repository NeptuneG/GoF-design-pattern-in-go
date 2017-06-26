package strategy

// Strategy
type strategy interface {
	callInterface() string
}

type strategyA struct {
}

type strategyB struct {
}

func (sa *strategyA) callInterface() string {
	return "call strategy A interface"
}

func (sb *strategyB) callInterface() string {
	return "call strategy A interface"
}

type context struct {
	strategy strategy
}

func (c *context) doAction() string {
	return c.strategy.callInterface()
}