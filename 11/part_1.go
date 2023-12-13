package main

import (
	"bufio"
	"fmt"
	"os"
)

func addColumn(starMap[][]string, colIndex int) [][]string {
	var returnMap [][]string
	for _, row := range starMap {
		newRow := row[:colIndex]
		newRow = append(newRow, ".")
		newRow = append(newRow, row[colIndex:]...)
		returnMap = append(returnMap, newRow)
	}
	return returnMap
}

func addRow(starMap [][]string, rowIndex int) [][]string {
	newRow := make([]string, len(starMap[0]))
	for i := 0; i < len(newRow); i++ {
		newRow[i] = "."
	}
	var returnMap [][]string
	returnMap = append(returnMap, starMap[:rowIndex]...)
	returnMap = append(returnMap, newRow)
	returnMap = append(returnMap, starMap[rowIndex:]...)
	return returnMap
}

func expand(starMap [][]string) [][]string {
	rowIndex := 0
	for rowIndex < len(starMap) {
		noStars := true
		for _, char := range starMap[rowIndex] {
			if char == "#" {
				noStars = false
			}
		}
		if noStars {
			starMap = addRow(starMap, rowIndex)
			rowIndex++
		}
		rowIndex++
	}

	colIndex := 0
	for colIndex < len(starMap[0]) {
		noStars := true
		for rowIndex := 0; rowIndex < len(starMap); rowIndex++ {
			if starMap[rowIndex][colIndex] == "#" {
				noStars = false
			}
		}
		if noStars {
			fmt.Println("Add column ", colIndex)
			starMap = addColumn(starMap, colIndex)
			colIndex++
		}
		colIndex++
	}

	return starMap
}

func getStars(starMap [][]string) [][]int {
	var stars [][]int
	for rowIndex, row := range(starMap) {
		for colIndex, char := range(row) {
			if char == "#" {
				stars = append(stars, []int{rowIndex, colIndex})
			}
		}
	}
	return stars
}

func parseInput() [][]string {
	var starMap [][]string
	scanner := bufio.NewScanner(os.Stdin)
	rowIndex := 0
	for scanner.Scan() {
		var row []string
		for _, rune := range scanner.Text() {
			char := string(rune)
			row = append(row, char)
		}
		starMap = append(starMap, row)
		rowIndex++
	}
	return starMap
}

func prettyPrintMatrix(matrix [][]string) {
	for _, row := range matrix {
		for _, char := range(row) {
			fmt.Printf("%v", char)
		}
		fmt.Println("")
	}
	fmt.Println("\n")
}

func sumDistances(stars [][]int) int {
	sum := 0
	for _, a := range stars {
		for _, b := range stars {
			height := a[0] - b[0]
			if height < 0 {
				height = height * -1
			}
			width := a[1] - b[1]
			if width < 0 {
				width = width * -1
			}
			distance := height + width
			sum += distance
		}
	}
	return sum
}

func main() {
	starMap := parseInput()
	prettyPrintMatrix(starMap)
	starMap = expand(starMap)
	prettyPrintMatrix(starMap)
	stars := getStars(starMap)
	fmt.Println(stars)
	fmt.Println("Total distance is ", sumDistances(stars) / 2)
}
