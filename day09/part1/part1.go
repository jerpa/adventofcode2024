package part1

import (
	c "github.com/jerpa/adventofcode2024/common"
)

func Part1(data []string) int {
	sum := 0
	m := []int{}
	f := 0
	for i := 0; i < len(data[0]); i += 2 {
		l := c.GetInt(string(data[0][i]))
		for j := 0; j < l; j++ {
			m = append(m, f)
		}
		if i+1 >= len(data[0]) {
			break
		}

		s := c.GetInt(string(data[0][i+1]))
		for j := 0; j < s; j++ {
			m = append(m, -1)
		}
		f++
	}

	for i := 0; i < len(m); i++ {
		if m[i] != -1 {
			continue
		}
		for j := len(m) - 1; j > i; j-- {
			if m[j] != -1 {
				m[i] = m[j]
				m[j] = -1
				break
			}
		}
	}

	for i := 0; i < len(m); i++ {
		if m[i] == -1 {
			continue
		}
		sum += m[i] * i
	}

	return sum
}
