package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("inputs/day4.txt")

	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(contents), "\n")

	passports := []map[string]string{
		make(map[string]string),
	}

	i := 0

	// Parse passports
	for _, row := range rows {
		if row == "" {
			i++
			passports = append(passports, make(map[string]string))
			continue
		}

		fields := strings.Split(row, " ")

		for _, field := range fields {
			keyValue := strings.Split(field, ":")

			if len(keyValue) != 2 {
				panic("key value keyValue could not be parsed")
			}

			passports[i][keyValue[0]] = keyValue[1]
		}
	}

	// Part1: validate passports
	validPart1 := 0
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

outer:
	for _, passport := range passports {
		for _, field := range requiredFields {
			if passport[field] == "" {
				continue outer
			}
		}

		validPart1++
	}

	fmt.Printf("valid: %v\n", validPart1)

	validPart2 := 0
	// Part2: validate passports
outer2:
	for _, passport := range passports {
		for _, field := range requiredFields {
			if passport[field] == "" {
				continue outer2
			}

			// Check if the format is correct
			if !isValid(field, passport[field]) {
				continue outer2
			}
		}

		validPart2++
	}

	fmt.Printf("valid: %v\n", validPart2)
}

var eyeColorRegexp = regexp.MustCompile("^#[0-9a-f]{6}$")
var pinRegexp = regexp.MustCompile("^[0-9]{9}$")

func isValid(field, value string) bool {
	switch field {
	case "byr":
		year, err := strconv.Atoi(value)

		return err == nil && (year >= 1920 && year <= 2002)
	case "iyr":
		year, err := strconv.Atoi(value)

		return err == nil && (year >= 2010 && year <= 2020)
	case "eyr":
		year, err := strconv.Atoi(value)

		return err == nil && (year >= 2020 && year <= 2030)
	case "hgt":
		if strings.HasSuffix(value, "cm") {
			height, err := strconv.Atoi(strings.TrimSuffix(value, "cm"))

			return err == nil && (height >= 150 && height <= 193)
		} else if strings.HasSuffix(value, "in") {
			height, err := strconv.Atoi(strings.TrimSuffix(value, "in"))

			return err == nil && (height >= 59 && height <= 76)
		} else {
			return false
		}
	case "hcl":
		return eyeColorRegexp.MatchString(value)
	case "ecl":
		return value == "amb" || value == "blu" || value == "brn" || value == "gry" || value == "grn" || value == "hzl" || value == "oth"
	case "pid":
		return pinRegexp.MatchString(value)
	}

	panic("field not recognized")
}
