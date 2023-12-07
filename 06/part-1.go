package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func winners(time int, record_distance int, step int) []int {
	var winners []int
	for hold_time := 1; hold_time < time; hold_time += step {
		distance := (time - hold_time) * hold_time
		if distance > record_distance {
			winners = append(winners, distance)
		}
	}
	return winners
}

func stringFieldsToIntSlice(field_string string) []int {
	var ints []int
	int_strings := strings.Fields(field_string)
	for _, int_string := range int_strings {
		converted, _ := strconv.Atoi(int_string)
		ints = append(ints, converted)
	}
	return ints
}

func main() {
	var times []int
	var distances []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var slice_to_load *[]int
		if strings.Contains(line, "Time") {
			slice_to_load = &times
		} else if strings.Contains(line, "Distance") {
			slice_to_load = &distances
		}
		chunks := strings.Split(line, ":")
		*slice_to_load = stringFieldsToIntSlice(chunks[1])
	}
	fmt.Println(times)
	fmt.Println(distances)

	part1 := 1
	for i, time := range times {
		wins := winners(time, distances[i], 1)
		part1 = part1 * len(wins)
	}
	fmt.Printf("Part 1 solution: %i\n", part1)

	part2Wins := winners(56977875, 546192711311139, 1)
	fmt.Printf("Part 2 solution: %i\n", len(part2Wins))
}
