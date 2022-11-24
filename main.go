package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/lukebrobbs/advent-of-code-2022/challenges"
)

var day = flag.Int("day", 1, "Advent of code day to run")

type challenge func(string) string

var dayInputs = map[int]challenge{
	1: challenges.Day1,
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(fmt.Sprintf("inputs/day-%d.go", *day))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
