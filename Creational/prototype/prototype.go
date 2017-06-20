package prototype

// Prototype
type prototype interface {
	clone() prototype
	GetInfo() string
}

type concretePrototype struct {
	info string
}

func (c *concretePrototype) clone() prototype {
	return &concretePrototype{c.info}
}

func (c *concretePrototype) GetInfo() string {
	return c.info
}