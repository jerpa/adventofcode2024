package part2

import (
	"github.com/jerpa/adventofcode2024/day04/shared"
)

func Part2(data []string) int {
	sum := countXmas(data)
	data = shared.RotateStrings(data)
	sum += countXmas(data)
	data = shared.RotateStrings(data)
	sum += countXmas(data)
	data = shared.RotateStrings(data)
	sum += countXmas(data)

	return sum
}
func countXmas(data []string) int {
	sum := 0
	for i := 0; i < len(data)-2; i++ {
		for j := 0; j < len(data[0])-2; j++ {
			if data[i][j] == 'M' && data[i][j+2] == 'S' && data[i+1][j+1] == 'A' && data[i+2][j] == 'M' && data[i+2][j+2] == 'S' {
				sum++
			}
		}
	}
	return sum
}
