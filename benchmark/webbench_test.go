package gotest

import (
	"testing"
)

func Test_Division(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("test_1: func Division can't passed")
	} else {
		t.Log("test_1 passed")
	}
	
}

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}
