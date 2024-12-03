package part2

import (
	"github.com/jerpa/adventofcode2024/day02/shared"
	"strings"
)

func Part2(data []string) int {
	sum := 0
	for _, v := range data {

		s := strings.Split(v, " ")

		if shared.CheckLevel(s) {
			sum++
		} else {
			for i := 0; i <= len(s); i++ {
				a := []string{}
				for x := range s {
					if i != x {
						a = append(a, s[x])
					}
				}
				if shared.CheckLevel(a) {
					sum++
					break
				}
			}
		}
	}

	return sum
}
