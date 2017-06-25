package command

// Reciever
type reciever struct {
	info string
}

func (r *reciever) action() string {
	return "reciever.action(): " + r.info
}

// Command
type command interface {
	execute() string
}

// Concrete command
type concreteCommand struct {
	rev *reciever
}

func (c *concreteCommand) execute() string {
	return c.rev.action()
}

// Invoker
type invoker struct {
	cmd command
}

func (i *invoker) invoke() string {
	return i.cmd.execute()
}
