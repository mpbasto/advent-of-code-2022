package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/lukebrobbs/advent-of-code-2022/challenges"
)

var day = flag.Int("day", 1, "Advent of code day to run")

func main() {
	flag.Parse()
	challenges := challenges.DayChallenges
	bytes, err := ioutil.ReadFile(fmt.Sprintf("inputs/day-%d.txt", *day))
	if err != nil {
		panic(err)
	}
	fmt.Println(challenges[*day](string(bytes)))
}
