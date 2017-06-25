package interpreter

import "testing"
import "strings"

func TestInterpreter(t *testing.T) {

	context := &context{strings.Fields("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")}

	// a dumb abstract syntax tree
	ast := make([]abstractExpression, 0)

	ast = append(ast, &nontermialExpr{})
	ast = append(ast, &nontermialExpr{})
	ast = append(ast, &terminalExpr{})

	actual := ""
	expect := "NNT"
	for _, expression := range ast {
		actual += expression.interpret(context)
	}

	if expect != actual {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expect, actual)
	}
}
