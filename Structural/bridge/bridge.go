package bridge

// Abstraction
type abstraction interface {
	operation() string
}

// Abstraction Implement
type concreteAbstractionImpl struct {
}

type refinedAbstraction struct {
	impl *concreteAbstractionImpl
}

func (c *concreteAbstractionImpl) operation() string {
	return "concrete abstaction implement"
}

func (r *refinedAbstraction) operation() string {
	return r.impl.operation()
}

