package main

import (
	"bufio"
	"fmt"
	"os"
)

type Board struct {
	Pipes [][]string
	StartRow int
	StartCol int
}

func checkUp(board *[][]string,
	currRow,
	currCol,
	prevRow,
	prevCol int) bool {
	source := (*board)[currRow][currCol]
	if currRow == 0 ||
		prevRow < currRow ||
		source == "-" ||
		source == "7" ||
		source == "F" {
		return false
	}
	destination := (*board)[currRow - 1][currCol]
	if destination == "-" ||
		destination == "L" ||
		destination == "." ||
		destination == "J" {
		return false
	}
	return true
}

func checkRight(board *[][]string,
	currRow,
	currCol,
	prevRow,
	prevCol int) bool {
	source := (*board)[currRow][currCol]
	if currCol == len((*board)[0]) - 1 ||
		prevCol > currCol ||
		source == "|" ||
		source == "J" ||
		source == "7" {
		return false
	}
	destination := (*board)[currRow][currCol + 1]
	if destination == "|" ||
		destination == "L" ||
		destination == "." ||
		destination == "F" {
		return false
	}
	return true
}

func checkDown(board *[][]string,
	currRow,
	currCol,
	prevRow,
	prevCol int) bool {
	source := (*board)[currRow][currCol]
	if currRow == len((*board)) - 1 ||
		prevRow > currRow ||
		source == "-" ||
		source == "L" ||
		source == "J" {
		return false
	}
	destination := (*board)[currRow + 1][currCol]
	if destination == "-" ||
		destination == "7" ||
		destination == "." ||
		destination == "F" {
		return false
	}
	return true
}

func checkLeft(board *[][]string,
	currRow,
	currCol,
	prevRow,
	prevCol int) bool {
	source := (*board)[currRow][currCol]
	if currCol == 0 ||
		prevCol < currCol ||
		source == "|" ||
		source == "L" ||
		source == "F" {
		return false
	}
	destination := (*board)[currRow][currCol - 1]
	if destination == "|" ||
		destination == "J" ||
		destination == "." ||
		destination == "7" {
		return false
	}
	return true
}

func parseInput() Board {
	var board Board
	var pipes [][]string
	scanner := bufio.NewScanner(os.Stdin)
	rowIndex := 0
	for scanner.Scan() {
		var row []string
		for colIndex, rune := range scanner.Text() {
			char := string(rune)
			if char == "S" {
				board.StartRow = rowIndex
				board.StartCol = colIndex
			}
			row = append(row, char)
		}
		pipes = append(pipes, row)
		rowIndex++
	}
	board.Pipes = pipes
	return board
}

func traverse(board Board) int {
	// colorReset := "\033[0m"
	// colorRed := "\033[31m"
	// colorGreen := "\033[32m"
	prettyMap := map[string]string{
		"|": "║",
		"L": "╚",
		"J": "╝",
		"7": "╗",
		"F": "╔",
		"-": "═",
		"": "\033[31m.\033[0m",
		"S": "\033[32m╬\033[0m",
	}
	currRow := board.StartRow
	currCol := board.StartCol
	prevRow := currRow
	prevCol := currCol
	steps := 0
	var viz [][]string
	numRows := len(board.Pipes)
	numCols := len(board.Pipes[0])
	for i := 0; i < numRows; i++ {
		viz = append(viz, make([]string, numCols))
	}
	for steps == 0 || currRow != board.StartRow || currCol != board.StartCol {
		viz[currRow][currCol] = board.Pipes[currRow][currCol]
		if steps > 0 && steps % 3500 == 0 {
			for _, row := range viz {
				for _, char := range(row) {
					fmt.Printf("%v", prettyMap[char])
				}
				fmt.Println("")
			}
			fmt.Println("\n")
		}
		tmpRow := currRow
		tmpCol := currCol
		if checkUp(&board.Pipes, currRow, currCol, prevRow, prevCol) {
			currRow--
		} else if checkRight(&board.Pipes, currRow, currCol, prevRow, prevCol) {
			currCol++
		} else if checkDown(&board.Pipes, currRow, currCol, prevRow, prevCol) {
			currRow++
		} else if checkLeft(&board.Pipes, currRow, currCol, prevRow, prevCol) {
			currCol--
		} else {
			panic(fmt.Sprintf("Nowhere to go currRow %v currCol %v prevRow %v prevCol %v", currRow, currCol, prevRow, prevCol))
		}
		prevRow = tmpRow
		prevCol = tmpCol
		steps++
	}
	for _, row := range viz {
		for _, char := range(row) {
			fmt.Printf("%v", prettyMap[char])
		}
		fmt.Println("")
	}
	fmt.Println("\n")
	return steps
}

func main() {
	board := parseInput()
	steps := traverse(board)
	fmt.Printf("Steps: %v, Farthest point: %v, Total tiles: %v\n", steps, steps/2, len(board.Pipes) * len(board.Pipes[0]))
}
