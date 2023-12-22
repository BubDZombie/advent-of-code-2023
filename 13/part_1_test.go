package main

import (
	"fmt"
	"testing"
)

func TestReflectHorizontal(t *testing.T) {
	input := []string{
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}
	var result int
	var expected int

	result = reflectHorizontal(input)
	expected = 400
	if result != expected {
		t.Errorf(
			fmt.Sprintf(
				"reflectHorizontal on %v expected %v got %v.",
				input,
				expected,
				result,
			),
		)
	}
	input = []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
	}
	result = reflectHorizontal(input)
	expected = 0
	if result != expected {
		t.Errorf(
			fmt.Sprintf(
				"reflectHorizontal on %v expected %v got %v.",
				input,
				expected,
				result,
			),
		)
	}
}
