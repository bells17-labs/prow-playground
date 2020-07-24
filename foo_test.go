package main

import (
	"testing"
)

func TestFoo(t *testing.T) {
	s := "bar"
	result := Foo(s)
	expect := s + " : foo"
	if result != expect {
		t.Errorf("Foo(%q) result is not %q: %q", s, expect, result)
	}
}
