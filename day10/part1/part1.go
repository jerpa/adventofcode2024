package part1

import (
	"fmt"
	c "github.com/jerpa/adventofcode2024/common"
)

func Part1(data []string) int {
	sum := 0
	for y := range data {
		for x := range data[y] {
			if data[y][x] == '0' {
				sum += testTrail(data, x, y, map[string]bool{})
			}
		}
	}

	return sum
}

func testTrail(data []string, x, y int, visited map[string]bool) int {
	if data[y][x] == '9' {
		pos := fmt.Sprintf("%d,%d", x, y)
		if _, ok := visited[pos]; ok {
			return 0
		}
		visited[pos] = true
		return 1
	}
	sum := 0
	val := c.GetInt(string(data[y][x]))

	if y > 0 && c.GetInt(string(data[y-1][x])) == val+1 {
		sum += testTrail(data, x, y-1, visited)
	}
	if x > 0 && c.GetInt(string(data[y][x-1])) == val+1 {
		sum += testTrail(data, x-1, y, visited)
	}
	if y < len(data)-1 && c.GetInt(string(data[y+1][x])) == val+1 {
		sum += testTrail(data, x, y+1, visited)
	}
	if x < len(data[y])-1 && c.GetInt(string(data[y][x+1])) == val+1 {
		sum += testTrail(data, x+1, y, visited)
	}
	return sum
}
