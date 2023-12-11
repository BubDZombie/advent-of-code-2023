package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringFieldsToIntSlice(field_string string) []int {
	var ints []int
	int_strings := strings.Fields(field_string)
	for _, int_string := range int_strings {
		converted, _ := strconv.Atoi(int_string)
		ints = append(ints, converted)
	}
	return ints
}

func getInputs() [][]int {
	var inputs [][]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputs = append(inputs, stringFieldsToIntSlice(scanner.Text()))
	}
	return inputs
}

func allZeroes(sequence []int) bool {
	for _, number := range(sequence) {
		if number != 0 {
			return false
		}
	}
	return true
}

func findNext(sequence []int) int {
	var pyramid [][]int
	pyramid = append(pyramid, sequence)
	prevRowIndex := 0
	for !allZeroes(pyramid[len(pyramid) - 1]) {
		var nextRow []int
		for i := 0; i <= len(pyramid[prevRowIndex]) - 2; i++ {
			nextRow = append(nextRow, pyramid[prevRowIndex][i + 1] - pyramid[prevRowIndex][i])
		}
		pyramid = append(pyramid, nextRow)
		prevRowIndex++
		fmt.Println(pyramid)
	}
	for i := len(pyramid) - 2; i >= 0; i-- {
		currRow := pyramid[i]
		nextRow := pyramid[i + 1]
		a := currRow[len(currRow) - 1]
		b := nextRow[len(nextRow) - 1]
		pyramid[i] = append(pyramid[i], a + b)
	}
	return(pyramid[0][len(pyramid[0]) - 1])
}

func main() {
	inputs := getInputs()
	sum := 0
	for i := 0; i < len(inputs); i++ {
		sum += findNext(inputs[i])
	}
	fmt.Println(sum)
}
