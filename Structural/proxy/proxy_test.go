package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	
	sub := &realSubject{}
	proxy := &proxy{sub}

	expect := "sub req"

	if proxy.request() != expect {
		t.Errorf("\nExpect:\t[ %s ]\nActual:\t[ %s ]", expect, proxy.request())
	}
	
}