package main

import (
	"fmt"
	"os"
	"strings"
)

var values = strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	panic("Value not found in index")
}

func convertValue(value string) int {
	return indexOf(value, values) + 1
}

func intersectBags(source1 string, source2 string) []string {
	intersections := []string{}
	for _, value := range strings.Split(source1, "") {
		if contains(strings.Split(source2, ""), value) {
			intersections = append(intersections, value)
		}
	}

	return intersections
}

func main() {
	data, err := os.ReadFile("./input")

	if err != nil {
		panic(err)
	}

	sum := 0
	rawBags := strings.Split(string(data), "\n")
	rawBagCount := len(rawBags) / 3
	for i := 0; i < rawBagCount; i++ {
		startIndex := i * 3
		rawBagSlice := rawBags[startIndex : startIndex+3]

		intersections := intersectBags(rawBagSlice[0], rawBagSlice[1])

		intersections = intersectBags(strings.Join(intersections, ""), rawBagSlice[2])

		sum += convertValue(intersections[0])
	}

	fmt.Println(sum)
}
