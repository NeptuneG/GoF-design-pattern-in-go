package iterator

// Iterator
type iterator interface {
	first()
	next()
	isDone() bool
	currentItem() interface{}
}

type concreteIterator struct {
	idx int
	aggregate aggregate
}

func (ci *concreteIterator) first() {
	ci.idx = 0
}

func (ci *concreteIterator) next() {
	if ci.idx < ci.aggregate.getSize() {
		ci.idx++
	}
}

func (ci *concreteIterator) isDone() bool {
	return ci.idx == ci.aggregate.getSize()
}

func (ci *concreteIterator) currentItem() interface{} {
	return ci.aggregate.getItem(ci.idx)
}