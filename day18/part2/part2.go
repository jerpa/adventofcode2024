package part2

import (
	"fmt"
	"github.com/RyanCarrier/dijkstra/v2"
)

func Part2(data []string, size, bytes int) int {
	m := map[string]struct{}{}
	for i := range data {
		m[data[i]] = struct{}{}
		if i < bytes {
			continue
		}
		v := map[string]int{}
		g := dijkstra.NewGraph()
		num := 0
		for y := 0; y <= size; y++ {
			for x := 0; x <= size; x++ {
				v[fmt.Sprintf("%d,%d", x, y)] = num
				g.AddEmptyVertex(num)
				num++
			}
		}
		for y := 0; y <= size; y++ {
			for x := 0; x <= size; x++ {
				if _, ok := m[fmt.Sprintf("%d,%d", x, y-1)]; !ok && y >= 0 {
					g.AddArc(v[fmt.Sprintf("%d,%d", x, y)], v[fmt.Sprintf("%d,%d", x, y-1)], 1)
				}
				if _, ok := m[fmt.Sprintf("%d,%d", x, y+1)]; !ok && y <= size {
					g.AddArc(v[fmt.Sprintf("%d,%d", x, y)], v[fmt.Sprintf("%d,%d", x, y+1)], 1)
				}
				if _, ok := m[fmt.Sprintf("%d,%d", x-1, y)]; !ok && x >= 0 {
					g.AddArc(v[fmt.Sprintf("%d,%d", x, y)], v[fmt.Sprintf("%d,%d", x-1, y)], 1)
				}
				if _, ok := m[fmt.Sprintf("%d,%d", x+1, y)]; !ok && x <= size {
					g.AddArc(v[fmt.Sprintf("%d,%d", x, y)], v[fmt.Sprintf("%d,%d", x+1, y)], 1)
				}
			}
		}
		_, err := g.Shortest(v["0,0"], v[fmt.Sprintf("%d,%d", size, size)])
		if err != nil {
			fmt.Println(data[i])
			return 0
		}
	}
	return -1
}
