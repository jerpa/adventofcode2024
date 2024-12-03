package main

import (
	c "github.com/jerpa/adventofcode2024/common"
	"github.com/jerpa/adventofcode2024/day01/part1"
	"github.com/jerpa/adventofcode2024/day01/part2"
	"time"
)

func main() {
	start := time.Now()

	c.Print("Part1: ", part1.Part1(c.GetData(Data)))
	c.Print("Took: ", time.Since(start).String())
	start = time.Now()
	c.Print("Part2: ", part2.Part2(c.GetData(Data)))
	c.Print("Took: ", time.Since(start).String())
}

var Data = `

`
