package part1

import (
	c "github.com/jerpa/adventofcode2024/common"
	"sort"
	"strings"
)

func Part1(data []string) int {
	sum := 0
	l := []int{}
	r := []int{}
	for _, v := range data {
		s := strings.Split(v, " ")
		l = append(l, c.GetInt(s[0]))
		r = append(r, c.GetInt(s[len(s)-1]))
	}
	sort.Ints(l)
	sort.Ints(r)
	for i := range l {
		diff := r[i] - l[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	return sum
}
