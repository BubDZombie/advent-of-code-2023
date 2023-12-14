package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	RecordString string
	Order []int
}

// Check to see if a record matches its order description.
// Return 0 for broken, 1 for potential match, 2 for exact match.
func check(record *string, order *[]int) int {
	adding := false
	orderIndex := 0
	springLength := 0
	partial := false
	for i := 0; i <= len(*record); i++ {
		var rune byte
		if i < len(*record) {
			rune = (*record)[i]
		} else {
			rune = 0
		}
		if !adding && rune == 35 {
			// 35 == # (American Keyboard)
			adding = true
			springLength = 1
		} else if adding && rune == 35 {
			springLength++
		} else if adding && rune != 35 {
			if orderIndex > len(*order) - 1 {
				return 1
			} else {
				if springLength != (*order)[orderIndex] {
					partial = true
				}
				if springLength > (*order)[orderIndex] {
					for orderIndex < len(*order) && springLength > (*order)[orderIndex] {
						if orderIndex == len(*order) - 1 && springLength > (*order)[orderIndex] {
							return 0
						}
						orderIndex++
					}
				}
			}
			orderIndex++
			adding = false
		}
	}
	if orderIndex == len(*order) && !partial {
		return 2
	}
	return 1
}

func getInputs() []Record {
	var records []Record
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		var record Record
		record.RecordString = fields[0]
		for _, intString := range strings.Split(fields[1], ",") {
			converted, _ := strconv.Atoi(intString)
			record.Order = append(record.Order, converted)
		}
		records = append(records, record)
	}
	return records
}

func permutations(record string, order []int) map[string]bool {
	var matches  = make(map[string]bool)
	for i, rune := range record {
		if rune == 63 {
			// 63 == ?
			newRecord := record[:i] + "#" + record[(i + 1):]
			result := check(&newRecord, &order)
			if result == 2 {
				matches[newRecord] = true
			} else if result == 1 {
				for key, _ := range permutations(newRecord, order) {
					matches[key] = true
				}
			}
		}
	}
	return matches
}

func main() {
	sum := 0
	for _, record := range getInputs() {
		fmt.Println(record)
		perms := permutations(record.RecordString, record.Order)
		fmt.Println(len(perms))
		sum += len(perms)
	}
	fmt.Println(sum)
}
