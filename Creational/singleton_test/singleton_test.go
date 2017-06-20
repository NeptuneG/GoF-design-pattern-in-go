package singleton_test

import (
	"testing"

	"github.com/NeptuneG/GoF-design-pattern-in-go/Creational/singleton"
)

func TestSingleton(t *testing.T) {

	instance1 := singleton.GetInstance()
	instance2 := singleton.GetInstance()

	if instance1 != instance2 {
		t.Errorf("expect instance1 equals to intance2 but not.")
	}
}
