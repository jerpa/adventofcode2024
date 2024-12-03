package part2

import (
	c "github.com/jerpa/adventofcode2024/common"
	"regexp"
)

func Part2(data []string) int {
	r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	sum := 0
	enabled := true
	for _, v := range data {
		s := r.FindAllStringSubmatch(v, -1)
		for _, x := range s {
			if x[0] == "do()" {
				enabled = true
			} else if x[0] == "don't()" {
				enabled = false
			} else if enabled {
				sum += c.GetInt(x[1]) * c.GetInt(x[2])
			}
		}

	}

	return sum
}
