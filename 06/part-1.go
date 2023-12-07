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

func main() {
	var times []int
	var distances []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Time") {
			time_chunks := strings.Split(line, ":")
			time_strings := strings.Fields(time_chunks[1])
			times = make([]int, len(time_strings))
			for i, time := range time_strings {
				time_int, _ := strconv.Atoi(time)
				times[i] = time_int
			}
		} else if strings.Contains(line, "Distance") {
			distance_chunks := strings.Split(line, ":")
			distance_strings := strings.Fields(distance_chunks[1])
			distances = make([]int, len(distance_strings))
			for i, distance := range distance_strings {
				distance_int, _ := strconv.Atoi(distance)
				distances[i] = distance_int
			}
		}
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
