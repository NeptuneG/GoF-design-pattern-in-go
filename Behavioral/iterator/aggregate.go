package iterator

// Aggregate
type aggregate interface {
	getItem(idx int) interface{}
	getSize() int
}

type concreteAggregate struct {
	objs []string
}

func (ca *concreteAggregate) getItem(idx int) interface{} {
	if idx < 0 || idx > len(ca.objs) {
		return nil
	}
	return ca.objs[idx]
}

func (ca *concreteAggregate) getSize() int {
	return len(ca.objs)
}