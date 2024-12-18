package part1

import (
	"fmt"
	"github.com/RyanCarrier/dijkstra/v2"
)

func Part1(data []string, size, bytes int) int {

	m := map[string]struct{}{}
	for i := 0; i < bytes; i++ {
		m[data[i]] = struct{}{}
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
	best, _ := g.Shortest(v["0,0"], v[fmt.Sprintf("%d,%d", size, size)])
	return int(best.Distance)
}
