package main

import (
	"sort"
)

type intSlice []int

func (s intSlice) Len() int           { return len(s) }
func (s intSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s intSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type SortAlgorithme func(a sort.Interface, left, k int)

func Quick(a sort.Interface, left, k int) {
	if left < k {
		p := partition(a, left, k)
		Quick(a, left, p-1)
		Quick(a, p+1, k)
	}
}

type intervales [][2]int

func Lamp(a sort.Interface, left, k int) {
	var (
		picked [2]int
		pivot  int
	)
	intervales := [][2]int{{left, k}}
	for len(intervales) > 0 {
		picked = intervales[0]
		if picked[1]-picked[0] > 0 {
			pivot = partition(a, picked[0], picked[1])
			intervales = append(intervales, [2]int{picked[0], pivot})
			intervales = append(intervales, [2]int{pivot + 1, picked[1]})
		}
		intervales = intervales[1:]
	}
}

func partition(a sort.Interface, left, right int) int {
	pivot := left
	a.Swap(pivot, right)
	index := left

	for i := left; i < right; i++ {
		if a.Less(i, right) {
			a.Swap(i, index)
			index += 1
		}
	}
	a.Swap(index, right)
	return index
}

func Sort(a intSlice, algo SortAlgorithme) intSlice {
	sorted := make([]int, len(a))
	copy(sorted, a)
	algo((intSlice)(sorted), 0, len(sorted)-1)
	return sorted
}
