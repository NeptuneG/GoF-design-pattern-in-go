package observer

type subject interface {
	setState(string)
	getState() string
	attach(observer)
	detach(observer)
	notify()
}

type concreteSubject struct {
	observers []observer
	state string
}

func (cs *concreteSubject) setState(state string) {
	cs.state = state
}

func (cs *concreteSubject) getState() string {
	return cs.state
}

func (cs *concreteSubject) attach(observer observer) {
	cs.observers = append(cs.observers, observer)
}

func (cs *concreteSubject) detach(observer observer) {
	for i, v := range cs.observers {
	    if v == observer {
        	cs.observers = append(cs.observers[:i], cs.observers[i+1:]...)
        	break
    	}
	}
}

func (cs *concreteSubject) notify() {
	for _, v := range cs.observers {
		v.update(cs)
	}
}