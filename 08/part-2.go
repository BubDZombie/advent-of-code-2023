package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"
)

type Node struct {
	Left string
	Right string
}

func advance(desertMap *map[string]Node, locations *[]string, locationIndex int, direction string, waitGroup *sync.WaitGroup) {
	defer (*waitGroup).Done()
	if direction == "L" {
		(*locations)[locationIndex] = (*desertMap)[(*locations)[locationIndex]].Left
	} else {
		(*locations)[locationIndex] = (*desertMap)[(*locations)[locationIndex]].Right
	}
}

func finished(locations *[]string) bool {
	for _, location := range *locations {
		if location[2] != 90 {
			return false
		}
	}
	return true
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

	steps := 0
	directionsIndex := 0
	var waitGroup sync.WaitGroup
	for !finished(&locations) {
		if directionsIndex > (len(directions) - 1) {
			directionsIndex = 0
		}
		direction := directions[directionsIndex]
		waitGroup.Add(len(locations))
		for i := 0; i < len(locations); i++ {
			go advance(&desertMap, &locations, i, direction, &waitGroup)
		}
		waitGroup.Wait()
		steps++
		directionsIndex++
	}
	fmt.Printf("It took %i steps to get from AAA to ZZZ.\n", steps)
}
