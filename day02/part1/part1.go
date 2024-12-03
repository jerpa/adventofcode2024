package part1

import (
	"github.com/jerpa/adventofcode2024/day02/shared"
	"strings"
)

func Part1(data []string) int {
	sum := 0
	for _, v := range data {

		s := strings.Split(v, " ")

		if shared.CheckLevel(s) {
			sum++
		}
	}

	return sum
}
