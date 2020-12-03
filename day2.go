package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("inputs/day2.txt")

	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(contents), "\n")

	matchedRowsPart1 := 0
	matchedRowsPart2 := 0

	for _, row := range rows {
		parts := strings.Split(row, " ")

		if len(parts) != 3 {
			panic("wrong format")
		}

		parts2 := strings.Split(parts[0], "-") // format "1-3"

		if len(parts2) != 2 {
			panic("wrong format")
		}

		min, _ := strconv.Atoi(parts2[0])
		max, _ := strconv.Atoi(parts2[1])

		char := rune(parts[1][0]) // format "x:"

		password := parts[2] // format "xyzxxxsdfsdf"

		count := 0

		// Part 1
		for _, c := range password {
			if c == char {
				count++
			}
		}

		if count >= min && count <= max {
			matchedRowsPart1++
		}

		// Part 2
		matchPosition1 := rune(password[min-1]) == char
		matchPosition2 := rune(password[max-1]) == char

		if (matchPosition1 && !matchPosition2) || (!matchPosition1 && matchPosition2) {
			matchedRowsPart2++
		}

	}

	println(matchedRowsPart1)
	println(matchedRowsPart2)
}
