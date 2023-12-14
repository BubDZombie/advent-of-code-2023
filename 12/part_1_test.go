package main

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	records := []string{"???.###", ".#?..#?...###.", "????.#####..######", ".#?..##...###."}
	orders := [][]int{{1, 1, 3}, {1, 1, 3}, {1, 5, 5}, {1, 1, 3}}
	outputs := []int{1, 2, 0, 0}
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
