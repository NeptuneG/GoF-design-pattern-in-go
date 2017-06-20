package factorymethod

// Product
type product interface {
	info() string
}

type concreteProduct struct {
}

func (c *concreteProduct) info() string {
	return "a concrete product"
}

// Factory
type factory interface {
	createProduct() product
}

type concreateFactory struct {
}

func (cf *concreateFactory) createProduct() product {
	return &concreteProduct{}
}
