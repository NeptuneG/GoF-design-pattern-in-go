package proxy

// Subject
type subject interface {
	request() string
}

type realSubject struct {
	
}

func (cs *realSubject) request() string {
	return "sub req"
}

// Proxy
type proxy struct {
	sub subject
}

func (p *proxy) request() string {
	p.checkSub()
	return p.sub.request()
}

func (p *proxy) checkSub() {
	if p.sub == nil {
		p.sub = &realSubject{}
	}
}