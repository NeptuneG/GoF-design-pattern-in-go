package chainofresposibility

// Handler interface
type handler interface {
	handleRequest(string) string
}

// Base handler
type baseHandler struct {
	info      string
	successor handler
}

func (bh *baseHandler) handleRequest(request string) string {
	baseRes := request + "->[BaseHandler]Default"
	if bh.successor != nil {
		return bh.successor.handleRequest(baseRes)
	}
	return baseRes
}

func (bh *baseHandler) setSuccessor(handler handler) {
	bh.successor = handler
}

func (bh *baseHandler) getSuccessor() handler {
	return bh.successor
}

// ConcreteHandlerA
type concreteHandlerA struct {
	*baseHandler
}

// ConcreteHandlerB
type concreteHandlerB struct {
	*baseHandler
	extraInfo string
}

func (cb *concreteHandlerB) handleRequest(request string) string {
	handlerBRes := request + "->[ConcHandlerB]" + cb.extraInfo
	if cb.successor != nil {
		return cb.successor.handleRequest(handlerBRes)
	}
	return handlerBRes
}
