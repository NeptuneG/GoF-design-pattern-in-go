package interpreter

// Context
type context struct {
	text []string
}

// Expression
type abstractExpression interface {
	interpret(*context) string
}

type terminalExpr struct {
}

func (te *terminalExpr) interpret(context *context) string {
	return "T"
}

type nontermialExpr struct {
}

func (ne *nontermialExpr) interpret(context *context) string {
	return "N"
}
