package part2

import c "github.com/jerpa/adventofcode2024/common"

func Part2(data []string) int {
	return 0
	sum := 0
	for _, v := range data {
		sum += c.GetInt(v)
	}

	return sum
}
