// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("./input")
	check(err)

    totalValueByElves := []int{};
	valuesByElves := strings.Split(string(data), "\n\n")
	for _, rawValues := range valuesByElves {
        totalValue := 0;
        for _, rawValue := range strings.Split(rawValues, "\n") {
            value, _ := strconv.Atoi(rawValue)
            totalValue += value
        }

        // totalValueByElves = totalValue
        totalValueByElves = append(totalValueByElves, totalValue)
	}

    sort.Sort(sort.Reverse(sort.IntSlice(totalValueByElves)))

    topThree := totalValueByElves[0:3];

    topThreeTotal := 0
    for _, topItem := range(topThree) {
        topThreeTotal += topItem
    }

    fmt.Println(topThreeTotal)
}
