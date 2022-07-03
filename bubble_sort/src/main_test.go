package main

import (
	"math/rand"
	"testing"
)

func TestBasic(t *testing.T) {
	var array []int = []int{3, 5, 8, 1, 4, 2, 7, 0, 6}
	var sorted []int = bubble_sort(array)
	for i := 0; i < len(sorted) - 1; i++ {
		if sorted[i] > sorted[i+1] {
			t.Errorf("sorted[%d] < sorted[%d]", i, sorted[i+1])
		} 
	}
}

func TestRandom1000(t *testing.T){
	var array []int = rand.Perm(1000)
	var sorted []int = bubble_sort(array)
	for i := 0; i < len(sorted) - 1; i++ {
		if sorted[i] > sorted[i+1] {
			t.Errorf("sorted[%d] != smaller then sorted[%d]", i, sorted[i+1])
		}
	}
}
