package iterator

import (
	"testing"
	"fmt"
	"strings"
)

func TestIterator(t *testing.T) {
	raw := strings.Fields("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
	ag := &concreteAggregate{raw}
	it := &concreteIterator{0, ag}

	for ; !it.isDone(); it.next() {
		fmt.Println(it.currentItem())
	}
}