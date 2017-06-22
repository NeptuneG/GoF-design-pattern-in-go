package flyweight

import (
	"testing"
)

func TestFlyweight(t *testing.T) {

	factory := &flyweightFactory{}

	if 0 != factory.getFactorySize() {
		t.Errorf("init flyweight factory error\n")
	}

	fw1 := factory.getFlyweight("foo")

	if 1 != factory.getFactorySize() || fw1.getState() != "foo" || fw1.operation() != "operation: foo" {
		t.Errorf("failed to get flyweight")
	}

	factory.getFlyweight("bar")
	if 2 != factory.getFactorySize() {
		t.Errorf("init flyweight factory error\n")
	}

	fw2 := factory.getFlyweight("foo")
	if 2 != factory.getFactorySize() || fw2.getState() != "foo" || fw2.operation() != "operation: foo" {
		t.Errorf("failed to get flyweight")
	}
}