package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Node struct {
	Left string
	Right string
}

// greatest common divisor via Euclidean algorithm
func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple via greatestCommonDivisor
func leastCommonMultiple(integers []int) int {
	a := integers[0]
	b := integers[1]
	integers = integers[2:]
	result := a * b / greatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = leastCommonMultiple([]int{result, integers[i]})
	}

	return result
}

func main() {
	var directions []string
	desertMap := make(map[string]Node)
	stdin := bufio.NewScanner(os.Stdin)
	directionsPattern := regexp.MustCompile(`^[LR]+$`)
	nodePattern := regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)
	for stdin.Scan() {
		line := stdin.Text()
		if directionsPattern.MatchString(line) {
			for _, rune := range(line) {
				directions = append(directions, string(rune))
			}
			fmt.Println(directions)
		} else if nodePattern.MatchString(line) {
			matches := nodePattern.FindAllStringSubmatch(line, -1)
			desertMap[matches[0][1]] = Node{matches[0][2], matches[0][3]}
		}
	}

	// Find the start locations.
	var locations []string
	startPattern := regexp.MustCompile(`..A`)
	fmt.Println(desertMap)
	for location, _ := range desertMap {
		if startPattern.MatchString(location) {
			locations = append(locations, location)
		}
	}
	fmt.Println(locations)

	// Routes are periodic, steps to first Z is the period.
	var periods []int
	for _, location := range locations {
		steps := 0
		directionsIndex := 0
		fmt.Printf("%s %i\n", location, steps)
		for location[2] != 90 {
			if directionsIndex > (len(directions) - 1) {
				directionsIndex = 0
			}
			direction := directions[directionsIndex]
			if direction == "L" {
				location = desertMap[location].Left
			} else {
				location = desertMap[location].Right
			}
			steps++
			directionsIndex++
		}
		periods = append(periods, steps)
	}
	fmt.Println(leastCommonMultiple(periods))
}
