package day1

import (
	"sort"
	"strconv"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ToInt(s string) int {
	result, err := strconv.Atoi(s)
	Check(err)
	return result
}

func Day1(input string, part2 bool) (x int) {
	food := strings.Split(string(input), "\n\n")
	elves := []int{}

	for _, elf := range food {
		sum := 0
		for _, line := range strings.Split(elf, "\n") {
			sum += ToInt(line)
		}
		elves = append(elves, sum)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	x = elves[0]
	if part2 {
		x += elves[1] + elves[2]
	}
	return
}
