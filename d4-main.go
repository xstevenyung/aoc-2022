package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type elf struct {
	min    int
	max    int
	values []int
}

func createElf(rawValue string) *elf {
	rawValues := strings.Split(rawValue, "-")

	min, _ := strconv.Atoi(rawValues[0])
	max, _ := strconv.Atoi(rawValues[1])

	values := []int{}

	for i := min; i <= max; i++ {
		values = append(values, i)
	}

	newElf := elf{min, max, values}

	return &newElf
}

func contains(s []int, value int) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}

	return false
}

func intersectValues(source1 []int, source2 []int) []int {
	intersections := []int{}
	for _, value := range source1 {
		if contains(source2, value) {
			intersections = append(intersections, value)
		}
	}

	return intersections
}

func main() {
	data, err := os.ReadFile("./input")

	if err != nil {
		panic("File not found")
	}

	count := 0
	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		elves := strings.Split(line, ",")

		firstElf := createElf(elves[0])
		secondElf := createElf(elves[1])

		intersections := intersectValues(firstElf.values, secondElf.values)

		if len(intersections) > 0 {
			count += 1
			continue
		}
	}

	fmt.Println(count)
}
