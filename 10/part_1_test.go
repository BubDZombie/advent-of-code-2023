package main

import (
	"testing"
)

func getTestBoard() [][]string {
	var board [][]string
	board = append(board, []string{"F", "-", "7", ".", "."})
	board = append(board, []string{"|", ".", "S", ".", "."})
	board = append(board, []string{"|", ".", "L", "-", "7"})
	board = append(board, []string{"|", ".", ".", ".", "|"})
	board = append(board, []string{"L", "-", "-", "-", "J"})
	return board
}

func TestCheckUp(t *testing.T) {
	board := getTestBoard()
	if checkUp(&board, 0, 0, 0, 0) {
		t.Errorf("Can't go up from the top row.")
	}
	if checkUp(&board, 4, 2, 4, 3) {
		t.Errorf("Can't go up from a dash.")
	}
	if checkUp(&board, 1, 0, 0, 0) {
		t.Errorf("Can't double back.")
	}
	if !checkUp(&board, 1, 2, 1, 2) {
		t.Errorf("Should be able to go up from start.")
	}
	if !checkUp(&board, 1, 0, 2, 0) {
		t.Errorf("Should be able to go up from pipe to pipe.")
	}
}

func TestCheckRight(t *testing.T) {
	board := getTestBoard()
	if checkRight(&board, 2, 4, 2, 4) {
		t.Errorf("Can't go right from the rightmost column.")
	}
	if !checkRight(&board, 4, 2, 4, 1) {
		t.Errorf("Should be able to go right from dash to dash.")
	}
}

func TestCheckDown(t *testing.T) {
	board := getTestBoard()
	if checkDown(&board, 4, 2, 4, 1) {
		t.Errorf("Can't go down from the bottom.")
	}
	if checkDown(&board, 2, 4, 3, 4) {
		t.Errorf("Can't double back.")
	}
	if !checkDown(&board, 2, 4, 2, 3) {
		t.Errorf("Should be able to go down from 7 to |.")
	}
}

func TestCheckLeft(t *testing.T) {
	board := getTestBoard()
	if checkLeft(&board, 0, 0, 0, 0) {
		t.Errorf("Can't go left from left edge.")
	}
	if !checkLeft(&board, 0, 1, 0, 2) {
		t.Errorf("Should be able to go left from - to F.")
	}
}
