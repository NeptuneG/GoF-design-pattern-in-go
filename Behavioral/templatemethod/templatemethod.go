package templatemethod

type operator interface {
	primitiveOpe1() string
	primitiveOpe2() string
}

type abstractOperator struct {
}

func (ao *abstractOperator) templateMethod(operator operator) string {
	return operator.primitiveOpe1() + "|" + operator.primitiveOpe2()
}

type concreteOpe1 struct {
	*abstractOperator
}

func (co1 *concreteOpe1) primitiveOpe1() string {
	return "Ope1.Ope1()"
}

func (co1 *concreteOpe1) primitiveOpe2() string {
	return "Ope1.Ope2()"
}

type concreteOpe2 struct {
	*abstractOperator
}

func (co2 *concreteOpe2) primitiveOpe1() string {
	return "Ope2.Ope1()"
}

func (co2 *concreteOpe2) primitiveOpe2() string {
	return "Ope2.Ope2()"
}