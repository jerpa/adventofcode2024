package shared

import "fmt"

func DoNothing() {
}

func AddAntinodes(antiNodes map[string]bool, data []string, max int) {
	for y := range data {
		for x := range data[y] {
			if data[y][x] != '.' {
				for yy := range data {
					for xx := range data[yy] {
						if y == yy && x == xx {
							continue
						}
						if data[yy][xx] == data[y][x] {
							deltaY := yy - y
							deltaX := xx - x
							if deltaY < 0 {
								deltaY = -deltaY
							}
							if deltaX < 0 {
								deltaX = -deltaX
							}
							if y+deltaY == yy {
								deltaY = -deltaY
							}
							if x+deltaX == xx {
								deltaX = -deltaX
							}
							for i := 1; i <= max; i++ {
								if y+(deltaY*i) < 0 || y+(deltaY*i) >= len(data) || x+(deltaX*i) < 0 || x+(deltaX*i) >= len(data[y]) {
									DoNothing()
								} else {
									antiNodes[fmt.Sprintf("%d,%d", x+(deltaX*i), y+(deltaY*i))] = true
								}
								if yy-(deltaY*i) < 0 || yy-(deltaY*i) >= len(data) || xx-(deltaX*i) < 0 || xx-(deltaX*i) >= len(data[y]) {
									DoNothing()
								} else {
									antiNodes[fmt.Sprintf("%d,%d", xx-(deltaX*i), yy-(deltaY*i))] = true
								}
							}
						}
					}
				}
			}
		}

	}
}
