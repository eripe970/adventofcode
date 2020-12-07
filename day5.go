package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	contents, err := ioutil.ReadFile("inputs/day5.txt")

	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(contents), "\n")

	highestSeat := int64(0)

	seen := make([]bool, 2048)

	for _, bits := range rows {

		if len(bits) > 10 {
			panic("wrong size of boarding card")
		}

		bits = strings.ReplaceAll(bits, "B", "1")
		bits = strings.ReplaceAll(bits, "F", "0")
		bits = strings.ReplaceAll(bits, "L", "0")
		bits = strings.ReplaceAll(bits, "R", "1")

		// This is the seat number since times 8 means moving 3 bits
		seat, _ := strconv.ParseInt(bits, 2, 0)

		if seat > highestSeat {
			highestSeat = seat
		}

		seen[seat] = true
	}

	println(highestSeat)

	startLooking := false

	for i := 0; i < len(seen); i++ {
		if !startLooking && seen[i] {
			startLooking = true
			continue
		}

		if startLooking && !seen[i] {
			println(i)
			break
		}
	}

}
