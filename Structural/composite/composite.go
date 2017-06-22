package composite

type component interface {
	operation() string
	add(c component)
	remove(c component)
	getChild(index int) interface{}
}

type composite struct {
	name string
	components []component
}

func (c *composite) operation() string {
	ret := c.name
	for _, comp := range c.components {
		ret += "-" + comp.operation()
	}
	return ret
}

func (c *composite) add(comp component) {
	c.components = append(c.components, comp)
}

func (c *composite) remove(comp component) {
	for i, v := range c.components {
		if v == comp {
			if len(c.components) == 1 {
				c.components = make([]component, 0)
			} else if i == len(c.components) - 1 {
				c.components = c.components[:i-1]
			} else if i == 0 {
				c.components = c.components[1:]
			} else {
				c.components = append(c.components[:i-1], c.components[i+1:]...)
			}
		}
	}
}

func (c *composite) getChild(index int) interface{} {
	return c.components[index]
}

type leaf struct {
	name string
}

func (l *leaf) operation() string {
	return l.name
}

func (l *leaf) add(c component) {
	
}

func (l *leaf) remove(c component) {
	
}

func (l *leaf) getChild(index int) interface{} {
	return nil
}