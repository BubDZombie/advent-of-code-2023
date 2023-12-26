package main

import (
	"fmt"
	"testing"
)

func TestReflectHorizontal(t *testing.T) {
	var inputs [][]string
	var outputs []int

	inputs = append(inputs, []string{
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	})
	outputs = append(outputs, 400)

	inputs = append(inputs, []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
	})
	outputs = append(outputs,  0)

	inputs = append(inputs, []string{
		"..#..#.#.##.#.#",
		"..#..#.#.##.#.#",
		".#..#.#..##..#.",
		"...###.######.#",
		".##............",
		"##.####......##",
		".###.#..#..##.#",
	})
	outputs = append(outputs, 100)

	for i, input := range inputs {
		result := reflectHorizontal(input)
		expected := outputs[i]
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
}

func TestReflectVertical(t *testing.T) {
	var inputs [][]string
	var outputs []int

	inputs = append(inputs, []string{
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	})
	outputs = append(outputs, 0)

	inputs = append(inputs, []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
	})
	outputs = append(outputs,  5)

	for i, input := range inputs {
		result := reflectVertical(input)
		expected := outputs[i]
		if result != expected {
			t.Errorf(
				fmt.Sprintf(
					"reflectVertical on %v expected %v got %v.",
					input,
					expected,
					result,
				),
			)
		}
	}
}
