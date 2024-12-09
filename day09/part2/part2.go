package part2

import (
	c "github.com/jerpa/adventofcode2024/common"
)

type block struct {
	length int
	data   int
}

func Part2(data []string) int {
	sum := 0
	m := []block{}
	f := 0
	for i := 0; i < len(data[0]); i += 2 {
		l := c.GetInt(string(data[0][i]))
		m = append(m, block{length: l, data: f})
		if i+1 >= len(data[0]) {
			break
		}

		s := c.GetInt(string(data[0][i+1]))
		m = append(m, block{length: s, data: -1})
		f++
	}

	for ; f >= 0; f-- {
	outer:
		for i := 0; i < len(m); i++ {
			if m[i].data == f {
				for j := 0; j < i; j++ {
					if m[j].data == -1 && m[j].length >= m[i].length {
						space := block{
							length: m[j].length - m[i].length,
							data:   -1,
						}
						m[j] = m[i]
						m[i].data = -1
						if m[i-1].data == -1 { // Merge
							m[i-1].length += m[i].length
							m = append(m[:i], m[i+1:]...)
						}
						if space.length > 0 {
							m = append(m[:j+1], append([]block{space}, m[j+1:]...)...)
						}
						break outer
					}
				}

			}
		}
	}
	pos := 0
	for i := range m {
		if m[i].data == -1 {
			pos += m[i].length
			continue
		}
		for j := 0; j < m[i].length; j++ {
			sum += pos * m[i].data
			pos++
		}
	}

	return sum
}
