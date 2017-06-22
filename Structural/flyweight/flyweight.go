package flyweight

// Flyweight
type flyweight interface {
	getState() string
	operation() string
}

type concreteFlyweight struct {
	state string
}

func (cf *concreteFlyweight) getState() string {
	return cf.state
}

func (cf *concreteFlyweight) operation() string {
	return "operation: " + cf.state
}

// Flyweight Factory
type flyweightFactory struct {
	flyweights []flyweight
}

func (ff *flyweightFactory) getFlyweight(key string) flyweight {
	for _, fw := range ff.flyweights {
		if key == fw.getState() {
			return fw
		}
	}

	cf := &concreteFlyweight{key}
	ff.flyweights = append(ff.flyweights, cf)
	return cf
}

func (ff *flyweightFactory) getFactorySize() int {
	return len(ff.flyweights)
}