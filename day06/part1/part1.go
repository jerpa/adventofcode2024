package part1

import (
	"fmt"
	c "github.com/jerpa/adventofcode2024/common"
)

func Part1(data []string) int {
	x, y := 0, 0
	for i := range data {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == '^' {
				x = j
				y = i
				break
			}
		}
	}
	pos := map[string]bool{}
	dir := c.Up
	for {
		pos[fmt.Sprintf("%d,%d", x, y)] = true
		switch dir {
		case c.Up:
			if y == 0 {
				return len(pos)
			}
			if data[y-1][x] == '#' {
				dir = c.Right
			} else {
				y--
			}
		case c.Right:
			if x == len(data[y])-1 {
				return len(pos)
			}
			if data[y][x+1] == '#' {
				dir = c.Down
			} else {
				x++
			}
		case c.Down:
			if y == len(data)-1 {
				return len(pos)
			}
			if data[y+1][x] == '#' {
				dir = c.Left
			} else {
				y++
			}
		case c.Left:
			if x == 0 {
				return len(pos)
			}
			if data[y][x-1] == '#' {
				dir = c.Up
			} else {
				x--
			}
		}
	}

}
