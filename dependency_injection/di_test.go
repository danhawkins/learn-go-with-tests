package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Danny")

	got := buffer.String()
	want := "Hello, Danny"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
