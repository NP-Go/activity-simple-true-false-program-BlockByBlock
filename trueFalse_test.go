package main

import (
	"testing"
)

func TestTrueFalse(t *testing.T) {
	got := compare(0)
	res := "Smaller than secret value"
	if got != res {
		t.Errorf("compare(0) = %s; want %s got %s", got, res, got)
	}

	got = compare(90)
	res = "Bigger than secret value"
	if got != res {
		t.Errorf("compare(90) = %s; want %s got %s", got, res, got)
	}

	got = compare(88)
	res = "Same as secret value"
	if got != res {
		t.Errorf("compare(88) = %s; want %s got %s", got, res, got)
	}

}
