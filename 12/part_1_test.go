package main

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	records := []string{"???.###", ".#?..#?...###.", "????.#####..######", ".#?..##...###.", "..?????#?##???."}
	orders := [][]int{{1, 1, 3}, {1, 1, 3}, {1, 5, 5}, {1, 1, 3}, {9, 1}}
	outputs := []int{1, 2, 0, 0, 1}
	for i := 0; i < len(records); i++ {
		response := check(&records[i], &orders[i])
		if  response!= outputs[i] {
			t.Errorf(
				fmt.Sprintf(
					"%v %v expected %v got %v.\n",
					records[i],
					orders[i],
					outputs[i],
					response))
		}
	}
}

func TestPermutations(t *testing.T) {
	perms := permutations(".??..??...?##.", []int{1,1,3})
	if len(perms) != 4 {
		t.Errorf(fmt.Sprintf(".??..??...?##. 1,1,3 got %v permutations, expected 4.\n", len(perms)))
	}
}
