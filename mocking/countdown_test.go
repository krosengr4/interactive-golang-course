package main

import (
	"testing"
	"bytes"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	act := buffer.String()
	exp := `3
2
1
Go!
`

	if act != exp {
		t.Errorf("ERROR: actual - %q expected - %q", act, exp)
	}

}



