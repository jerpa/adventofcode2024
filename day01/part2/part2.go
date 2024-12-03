package part2

import (
	c "github.com/jerpa/adventofcode2024/common"
	"strings"
)

func Part2(data []string) int {
	sum := 0
	l := []int{}
	r := []int{}
	for _, v := range data {
		s := strings.Split(v, " ")
		l = append(l, c.GetInt(s[0]))
		r = append(r, c.GetInt(s[len(s)-1]))
	}

	for i := range l {
		num := 0
		for _, v := range r {
			if l[i] == v {
				num++
			}
		}
		sum += num * l[i]
	}

	return sum
}
