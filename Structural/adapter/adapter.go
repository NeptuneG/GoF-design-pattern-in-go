package adapter

type target interface {
	request() string
}

type adaptee struct {
}

func (a *adaptee) specificRequest() string {
	return "adaptee's request & response"
}

type adapter struct {
	adaptee *adaptee
}

func (a *adapter) request() string {
	return a.adaptee.specificRequest()
}
