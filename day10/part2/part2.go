package part2

import (
	c "github.com/jerpa/adventofcode2024/common"
)

func Part2(data []string) int {
	sum := 0
	for y := range data {
		for x := range data[y] {
			if data[y][x] == '0' {
				sum += testTrail(data, x, y)
			}
		}
	}

	return sum
}

func testTrail(data []string, x, y int) int {
	if data[y][x] == '9' {
		return 1
	}
	sum := 0
	val := c.GetInt(string(data[y][x]))

	if y > 0 && c.GetInt(string(data[y-1][x])) == val+1 {
		sum += testTrail(data, x, y-1)
	}
	if x > 0 && c.GetInt(string(data[y][x-1])) == val+1 {
		sum += testTrail(data, x-1, y)
	}
	if y < len(data)-1 && c.GetInt(string(data[y+1][x])) == val+1 {
		sum += testTrail(data, x, y+1)
	}
	if x < len(data[y])-1 && c.GetInt(string(data[y][x+1])) == val+1 {
		sum += testTrail(data, x+1, y)
	}
	return sum
}
