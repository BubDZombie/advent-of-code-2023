package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInputs() [][]string {
	var records [][]string
	var record []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			records = append(records, record)
			record = nil
		} else {
			record = append(record, line)
		}
	}
	records = append(records, record)
	return records
}

func reflectHorizontal(record []string) int {
	for top := 0; top < len(record) - 1; top++ {
		bottom := top + 1
		i := top
		j := bottom
		mismatched := 0
		for i >= 0 && j < len(record) {
			for k, _ := range record[i] {
				if record[i][k] != record[j][k] {
				mismatched++
				}
			}
			i--
			j++
		}
		if mismatched > 1 {
			continue
		} else if (i == -1 || j == len(record)) && mismatched == 1 {
			return 100 * bottom
		}
	}
	return 0
}

func reflectVertical(record []string) int {
	for left := 0; left < len(record[0]) - 1; left++ {
		right := left + 1
		mismatched := 0
		i := left
		j := right
		for mismatched < 2 && i >= 0 && j < len(record[0]) {
			for _, line := range record {
				if line[i] != line[j] {
					mismatched++
				}
			}
			if mismatched > 1 {
				continue
			} else if (i == 0 || j == len(record[0]) - 1) && mismatched == 1 {
				return right
			}
			i--
			j++
		}
	}
	return 0
}

func main() {
	sum := 0
	for _, record := range getInputs() {
		for _, line := range record {
			fmt.Println(line)
		}
		response := reflectHorizontal(record)
		fmt.Println("reflectHorizontal", response)
		sum += response
		response = reflectVertical(record)
		fmt.Println("reflectVertical", response)
		sum += response
	}
	fmt.Println(sum)
}
