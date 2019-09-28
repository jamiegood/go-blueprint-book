package trace

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// TestNew ...
func TestNew(t *testing.T) {
	//t.Error("We haven't written our tests yet")
	var buf bytes.Buffer
	fmt.Println("First Buf is:; ")
	fmt.Println(buf)
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from new should not be nil")
	} else {
		tracer.Trace("Hello trace package")
		if buf.String() != "Hello trace package\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())

		}

	}
	fmt.Println("Second Buf is:; ")
	fmt.Println(buf)
}
