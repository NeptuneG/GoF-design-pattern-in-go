package visitor

import (
	"fmt"
)

// Visitor
type visitor interface {
	visitElementA(element element) string
	visitElementB(element element) string
}

type concreteVisitorA struct {
}

func (cva *concreteVisitorA) visitElementA(element element) string {
	return "visitorA.visits.elementA"
}

func (cva *concreteVisitorA) visitElementB(element element) string {
	return "visitorA.visits.elementB"
}

type concreteVisitorB struct {
}

func (cvb *concreteVisitorB) visitElementA(element element) string {
	return "visitorB.visits.elementA"
}

func (cvb *concreteVisitorB) visitElementB(element element) string {
	return "visitorB.visits.elementB"
}

// Element
type element interface {
	accept(visitor visitor) string
}

type concreteElementA struct {
}

func (cea *concreteElementA) accept(visitor visitor) string {
	fmt.Println(visitor.visitElementA(cea))
	return "eleA is visited"
}

type concreteElementB struct {
}

func (ceb *concreteElementB) accept(visitor visitor) string {
	fmt.Println(visitor.visitElementB(ceb))
	return "eleB is visited"
}