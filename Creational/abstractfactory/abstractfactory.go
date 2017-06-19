package abstractfactory

// Product
type abstractProduct interface {
	info() string
}

type productA struct {
}

type productB struct {
}

func (a *productA) info() string {
	return "Product A"
}

func (b *productB) info() string {
	return "Product B"
}

// Factory
type abstractFactory interface {
	// X createProductA() *abstractProduct
	// X createProductB() *abstractProduct
	// https://stackoverflow.com/questions/13511203/why-cant-i-assign-a-struct-to-an-interface
	createProductA() abstractProduct
	createProductB() abstractProduct
}

type concreateFactory struct {
}

func (cf *concreateFactory) createProductA() abstractProduct {
	return &productA{}
}

func (cf *concreateFactory) createProductB() abstractProduct {
	return &productB{}
}
