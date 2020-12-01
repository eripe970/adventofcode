package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	contents, err := ioutil.ReadFile("inputs/day1.txt")

	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(contents), "\n")

	var numbers []int

	for _, s := range rows {
		n, err := strconv.Atoi(s)

		if err != nil {
			panic(err)
		}

		numbers = append(numbers, n)
	}

	println("part a, not a beauty but works: ")

	for i, _ := range numbers {
		for j, _ := range numbers {
			if i > j {
				if numbers[i]+numbers[j] == 2020 {
					fmt.Printf("sum: %d\n", numbers[i]*numbers[j])
				}
			}
		}
	}

	println("part b, not a beauty but works: ")

	for i, _ := range numbers {
		for j, _ := range numbers {
			if j > i {
				break
			}

			for k, _ := range numbers {
				if k > j {
					break
				}

				sumAll := numbers[i] + numbers[j] + numbers[k]
				if sumAll == 2020 {
					fmt.Printf("sum: %d\n", numbers[i]*numbers[j]*numbers[k])
				}
			}

		}
	}
}
