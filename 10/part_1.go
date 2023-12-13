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
	Steps int
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

// Enhance a set of pipes https://www.youtube.com/watch?v=KiqkclCJsZs
// so that we can edge find and keep our sanity.
//
//           .O.
// J becomes OO.
//           ...
//
//           ...
// . becomes .!.
//           ...
func enhance(pipes [][]string) [][]string {
	charMap := map[string][][]string{
		"J": {
			{".", "║", "."},
			{"═", "╝", "."},
			{".", ".", "."}},
		"L": {
			{".", "║", "."},
			{".", "╚", "═"},
			{".", ".", "."}},
		"F": {
			{".", ".", "."},
			{".", "╔", "═"},
			{".", "║", "."}},
		"7": {
			{".", ".", "."},
			{"═", "╗", "."},
			{".", "║", "."}},
		"|": {
			{".", "║", "."},
			{".", "║", "."},
			{".", "║", "."}},
		"-": {
			{".", ".", "."},
			{"═", "═", "═"},
			{".", ".", "."}},
		"": {
			{".", ".", "."},
			{".", "!", "."},
			{".", ".", "."}},
		"S": {
			{".", "║", "."},
			{"═", "╬", "═"},
			{".", "║", "."}},
	}
	numCols := len(pipes[0])
	var enhanced [][]string
	for smallRowIndex, row := range pipes {
		for i := 0; i < 3; i++ {
			enhanced = append(enhanced, make([]string, 3 * numCols))
		}
		for smallColIndex, smallChar := range row {
			for bigMapRowIndex, bigMapRow := range charMap[smallChar] {
				for bigMapColIndex, bigMapChar := range bigMapRow {
					enhanced[3 * smallRowIndex + bigMapRowIndex][3 * smallColIndex + bigMapColIndex] = bigMapChar
				}
			}
		}
	}
	return enhanced
}

// Start at 0,0 of an enhanced board. Spread over empty tiles until you can't
// spread any more. Count remaining empty tiles.
// Return int count of remaining tiles.
func ooze(enhanced [][]string) int {
	numRows := len(enhanced)
	numCols := len(enhanced[0])
	to_process := [][]int{{0, 0}}
	iterations := 0
	var address []int
	for len(to_process) > 0 {
		address, to_process = to_process[0], to_process[1:]
		row, col := address[0], address[1]
		if enhanced[row][col] == "." || enhanced[row][col] == "!" {
			enhanced[row][col] = "░"
		}
		for _, look := range [][]int{{row - 1, col}, {row, col + 1}, {row + 1, col}, {row, col - 1}} {
			if look[0] >= 0 && look[0] < numRows && look[1] >= 0 && look[1] < numCols &&
			(enhanced[look[0]][look[1]] == "." || enhanced[look[0]][look[1]] == "!") {
				to_process = append([][]int{look}, to_process...)
			}
		}
		iterations++
	}

	// Count !s
	bangs := 0
	for _, row := range enhanced {
		for _, char := range row {
			if char == "!" {
				bangs++
			}
		}
	}
	return bangs
}

func prettyPrintMatrix(pipes [][]string, color bool) {
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
		"!": "\033[31m!\033[0m",
		"": ".",
		"S": "\033[32m╬\033[0m",
	}
	for _, row := range pipes {
		for _, char := range(row) {
			var pretty string
			pretty, ok := prettyMap[char]
			if !ok || !color {
				pretty = char
			}
			fmt.Printf("%v", pretty)
		}
		fmt.Println("")
	}
	fmt.Println("\n")
}

func traverse(board Board) Board {
	currRow := board.StartRow
	currCol := board.StartCol
	prevRow := currRow
	prevCol := currCol
	steps := 0
	var cleanPipes [][]string
	numRows := len(board.Pipes)
	numCols := len(board.Pipes[0])
	for i := 0; i < numRows; i++ {
		cleanPipes = append(cleanPipes, make([]string, numCols))
	}
	for steps == 0 || currRow != board.StartRow || currCol != board.StartCol {
		cleanPipes[currRow][currCol] = board.Pipes[currRow][currCol]
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
	var cleanBoard Board
	cleanBoard.Pipes = cleanPipes
	cleanBoard.StartRow = board.StartRow
	cleanBoard.StartCol = board.StartCol
	cleanBoard.Steps = steps
	return cleanBoard
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

func main() {
	board := parseInput()
	board = traverse(board)
	steps := board.Steps
	fmt.Printf("Steps: %v, Farthest point: %v, Total tiles: %v\n", steps, steps/2, len(board.Pipes) * len(board.Pipes[0]))
	board.Pipes = enhance(board.Pipes)
	enclosed := ooze(board.Pipes)
	fmt.Printf("The loop encloses %v tiles.\n", enclosed)
}
