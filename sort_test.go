package main

import (
	"testing"
)

type table [2][]int

var tables = [...]table{
	table{{},
		{}},
	table{{1},
		{1}},
	table{{2, 1},
		{1, 2}},
	table{{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
}

func sliceEq(a, b intSlice) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestQuick(t *testing.T) {
	for _, c := range tables {
		qsorted := Sort(c[0], Quick)
		expected := c[1]
		if !sliceEq(qsorted, expected) {
			t.Logf("quick: got %v expected %v", qsorted, expected)
			t.Fail()
		}
	}
}

func BenchmarkQuick(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range tables {
			Sort(c[0], Quick)
		}
	}
}

func TestLamp(t *testing.T) {
	for _, c := range tables {
		lsorted := Sort(c[0], Lamp)
		expected := c[1]
		if !sliceEq(lsorted, expected) {
			t.Logf("quick: got %v expected %v", lsorted, expected)
			t.Fail()
		}
	}
}

func BenchmarkLamp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range tables {
			Sort(c[0], Lamp)
		}
	}
}
