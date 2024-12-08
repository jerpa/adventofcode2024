package part1

import "github.com/jerpa/adventofcode2024/day08/shared"

func Part1(data []string) int {

	antiNodes := map[string]bool{}
	shared.AddAntinodes(antiNodes, data, 1)
	return len(antiNodes)

}
