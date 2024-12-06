package part2

import (
	"fmt"
	c "github.com/jerpa/adventofcode2024/common"
)

func Part2(data []string) int {
	sum := 0
	startX, startY := 0, 0
findPos:
	for i := range data {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == '^' {
				startX = j
				startY = i
				break findPos
			}
		}
	}
	for y := range data {
		for x := range data[y] {
			if data[y][x] == '#' || data[y][x] == '^' {
				continue
			}
			if isLoop(data, startX, startY, x, y) {
				sum++
			}
		}
	}
	return sum
}

func isLoop(data []string, x, y, obsX, obsY int) bool {
	pos := map[string]bool{}
	dir := c.Up

	for {
		p := fmt.Sprintf("%d,%d,%d", x, y, dir)
		if pos[p] {
			return true
		}
		pos[p] = true
		switch dir {
		case c.Up:
			if y == 0 {
				return false
			}
			if data[y-1][x] == '#' || (y-1 == obsY && x == obsX) {
				dir = c.Right
			} else {
				y--
			}
		case c.Right:
			if x == len(data[y])-1 {
				return false
			}
			if data[y][x+1] == '#' || (y == obsY && x+1 == obsX) {
				dir = c.Down
			} else {
				x++
			}
		case c.Down:
			if y == len(data)-1 {
				return false
			}
			if data[y+1][x] == '#' || (y+1 == obsY && x == obsX) {
				dir = c.Left
			} else {
				y++
			}
		case c.Left:
			if x == 0 {
				return false
			}
			if data[y][x-1] == '#' || (y == obsY && x-1 == obsX) {
				dir = c.Up
			} else {
				x--
			}
		}
	}

}
