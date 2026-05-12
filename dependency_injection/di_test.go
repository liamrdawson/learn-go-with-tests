package main

import (
	"bytes"
	"testing"
)

/*
*	IMPORTANT
*
*	*	Give me the address of this variable (go from value -> pointer)
*	&	Give me the value at this address (go from pointer -> value)
*
*	&buffer		Take the address of buffer - produces *bytes.Buffer
*	*testing.T	t is a pointer to a testing.T struct - a type annotation
*	*t			Give me the vlaue that t points to - dereferencing
 */
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	// Because we're passing &buffer we're pointing at the address of
	// bytes.Buffer. Any writes inside Greet would affect the SAME buffer in
	// memory, not a copy. If we passed without & then Go would copy the struct,
	// Greet would write to the copy, and the original buffer would remain empty.
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
