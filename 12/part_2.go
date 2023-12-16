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
	//fmt.Printf("check %v %v\n", *record, *order)
	orderTotal := 0
	orderMax := 0
	for _, size := range *order {
		orderTotal += size
		if size > orderMax {
			orderMax = size
		}
	}
	springTotal := 0

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
			springTotal++
		} else if adding && rune == 35 {
			springLength++
			springTotal++
		} else if adding && rune != 35 {
			if springLength > orderMax {
				return 0
			} else if orderIndex > len(*order) - 1 || springLength != (*order)[orderIndex] {
				partial = true
			}
			orderIndex++
			adding = false
		}
		if springTotal > orderTotal {
			return 0
		}
	}
	//fmt.Printf("orderTotal %v springTotal %v\n", orderTotal, springTotal)
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
		record.RecordString = strings.Repeat(fields[0], 5)
		for i := 0; i < 5; i++ {
			for _, intString := range strings.Split(fields[1], ",") {
				converted, _ := strconv.Atoi(intString)
				record.Order = append(record.Order, converted)
			}
		}
		records = append(records, record)
	}
	return records
}

func permutations(record string, order []int) []string {
	var results = make(map[string]int)
	toProcess := []string{record}
	for len(toProcess) > 0 {
		processing := toProcess[0]
		toProcess = toProcess[1:]
		_, tried := results[processing]
		if !tried {
			results[processing] = check(&processing, &order)
			//fmt.Println(processing, results[processing])
			if results[processing] == 1 {
				for i, rune := range processing {
					if rune == 63 {
						// 63 == ?
						newRecord := processing[:i] + "#" + processing[(i + 1):]
						toProcess = append([]string{newRecord}, toProcess...)
					}
				}
			}
		}
	}
	var matches []string
	for r, result := range results {
		if result == 2 {
			matches = append(matches, r)
		}
	}
	return matches
}

func main() {
	sum := 0
	for _, record := range getInputs() {
		fmt.Println(record)
		perms := permutations(record.RecordString, record.Order)
		// for _, perm := range perms {
		// 	fmt.Println(perm)
		// }
		fmt.Println(len(perms))
		sum += len(perms)
	}
	fmt.Println(sum)
}
