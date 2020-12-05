package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("inputs/day3.txt")

	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(contents), "\n")

	trees := make([][]bool, len(rows))

	for i, row := range rows {
		trees[i] = make([]bool, len(row))

		for j, char := range row {
			trees[i][j] = char == '#'
		}
	}

	a := calcSlopes(1, 1, trees)
	b := calcSlopes(3, 1, trees) // Part 1
	c := calcSlopes(5, 1, trees)
	d := calcSlopes(7, 1, trees)
	e := calcSlopes(1, 2, trees)

	fmt.Printf("Part1: %d\n", b)
	fmt.Printf("Part2: %d\n", a*b*c*d*e)
}

func calcSlopes(right, down int, area [][]bool) int {
	trees := 0
	column := 0

	for i := 0; i*down < len(area); i++ {

		if area[i*down][column] {
			trees++
		}

		column = (column + right) % len(area[i])

	}

	return trees
}
