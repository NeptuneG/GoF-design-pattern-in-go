package facade

type subsys1 struct {
}

type subsys2 struct {
}

type facade struct {
	sys1 *subsys1
	sys2 *subsys2
}

func (s *subsys1) foo() string {
	return "foo"
}

func (s *subsys2) bar() string {
	return "bar"
}

func (f *facade) wrapper() string {
	return "wrapper: [" + f.sys1.foo() + f.sys2.bar() + "]"
}