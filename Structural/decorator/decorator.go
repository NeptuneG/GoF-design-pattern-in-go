package decorator

// Component
type component interface {
	operation() string
}

type concreteComponent struct {
	
}

func (c *concreteComponent) operation() string {
	return "dumb comp"
}

// Decorator
type decorator struct {
	component component
}

func (d *decorator) addedBehavior(raw string) string {
	return "###" + raw + "###"
}

func (d *decorator) operation() string {
	return d.addedBehavior(d.component.operation())
}


