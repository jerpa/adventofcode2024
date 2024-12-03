package part1

import (
	c "github.com/jerpa/adventofcode2024/common"
	"regexp"
)

func Part1(data []string) int {
	r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	sum := 0
	for _, v := range data {
		s := r.FindAllStringSubmatch(v, -1)
		for _, x := range s {
			sum += c.GetInt(x[1]) * c.GetInt(x[2])
		}

	}

	return sum
}
