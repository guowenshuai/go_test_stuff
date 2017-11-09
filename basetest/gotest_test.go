package gotest

import (
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("test_1: func Division can't passed")
	} else {
		t.Log("test_1 passed")
	}
}

func TestDivision(t *testing.T) {
	t.Error("test_2: can't passed")
}

