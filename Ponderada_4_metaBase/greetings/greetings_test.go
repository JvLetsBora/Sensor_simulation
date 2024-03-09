package greetings

import (
	"testing"
)

func TestAbs(t *testing.T) {
	// -1, 0, 1
	if Abs(-1) < 0 {
		t.Error("iijiji", -1)
	}

}
