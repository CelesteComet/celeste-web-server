package main

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	hello := SayHello()
	if hello != "HELLO" {
		t.Error("Expected SayHello to return 'HELLO'")
	}
}
