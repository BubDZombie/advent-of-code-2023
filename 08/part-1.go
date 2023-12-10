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

	steps := 0
	directionsIndex := 0
	location := "AAA"
	for location != "ZZZ" {
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
	fmt.Printf("It took %i steps to get from AAA to ZZZ.\n", steps)
}
