package goarg

import (
	"testing"
)

func TestReturn(t *testing.T) {
	var i int
	FlagList = []IFlag{}
	AddArg(&i, "int", 10, "int flag help", true)

	if len(FlagList) != 1 {
		t.Errorf("Expected length 1, got %d", len(FlagList))
	}
}
