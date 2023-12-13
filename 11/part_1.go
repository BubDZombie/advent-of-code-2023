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

func getEmptyColumns(starMap [][]string) []int {
	var columns []int
	colIndex := 0
	for colIndex < len(starMap[0]) {
		noStars := true
		for rowIndex := 0; rowIndex < len(starMap); rowIndex++ {
			if starMap[rowIndex][colIndex] == "#" {
				noStars = false
			}
		}
		if noStars {
			columns = append(columns, colIndex)
		}
		colIndex++
	}
	return columns
}

func getEmptyRows(starMap [][]string) []int {
	var rows []int
	for rowIndex, row := range starMap {
		noStars := true
		for _, char := range row {
			if char == "#" {
				noStars = false
			}
		}
		if noStars {
			rows = append(rows, rowIndex)
		}
	}
	return rows
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max (a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sumDistances(stars [][]int, emptyRows []int, emptyColumns []int, speedOfLight int) int {
	sum := 0
	for _, a := range stars {
		for _, b := range stars {
			height := a[0] - b[0]
			if height < 0 {
				height = height * -1
			}
			for _, row := range emptyRows {
				if row > min(a[0], b[0]) && row < max(a[0], b[0]) {
					height += speedOfLight - 1
				}
			}
			width := a[1] - b[1]
			if width < 0 {
				width = width * -1
			}
			for _, column := range emptyColumns {
				if column > min(a[1], b[1]) && column < max(a[1], b[1]) {
					width += speedOfLight - 1
				}
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
	emptyRows := getEmptyRows(starMap)
	emptyColumns := getEmptyColumns(starMap)
	fmt.Println("Empty rows: ", emptyRows, " empty columns: ", emptyColumns)
	stars := getStars(starMap)
	fmt.Println("Stars: ", stars)
	fmt.Println("Total distance for part 1 is ", sumDistances(stars, emptyRows, emptyColumns, 2) / 2)
	fmt.Println("Total distance with light speed 10 is ", sumDistances(stars, emptyRows, emptyColumns, 10) / 2)
	fmt.Println("Total distance with light speed 100 is ", sumDistances(stars, emptyRows, emptyColumns, 100) / 2)
	fmt.Println("Total distance for part 2 is ", sumDistances(stars, emptyRows, emptyColumns, 1000000) / 2)
}
