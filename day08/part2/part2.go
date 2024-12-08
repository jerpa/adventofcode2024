package part2

import (
	"fmt"
	"github.com/jerpa/adventofcode2024/day08/shared"
)

func Part2(data []string) int {

	antiNodes := map[string]bool{}
	shared.AddAntinodes(antiNodes, data, len(data))
	sum := 0
	for y := range data {
		for x := range data[y] {
			if _, ok := antiNodes[fmt.Sprintf("%d,%d", x, y)]; ok || data[y][x] != '.' {
				sum++
			}
		}
	}
	return sum
}
