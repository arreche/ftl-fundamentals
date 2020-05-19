package hello_test

import (
	"hello"
	"testing"
)

func TestReturnGreeting(t *testing.T) {
	got := hello.ReturnGreeting("Hi there")
	want := "Hi there yourself!"
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
