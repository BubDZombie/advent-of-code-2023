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
		mismatched := false
		for i >= 0 && j < len(record) {
			if record[i] != record[j] {
				mismatched = true
				break
			}
			i--
			j++
		}
		if mismatched {
			continue
		} else if i == 0 || j == len(record) - 1 {
			return 100 * bottom
		}
	}
	return 0
}

func main() {
	records := getInputs()
	fmt.Println(records)
}
