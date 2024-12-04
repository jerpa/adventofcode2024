package part1

import (
	"github.com/jerpa/adventofcode2024/day04/shared"
	"strings"
)

func Part1(data []string) int {
	sum := 0
	for _, v := range data {
		sum += strings.Count(v, "XMAS")
		sum += strings.Count(v, "SAMX")
	}
	data = shared.RotateStrings(data)
	for _, v := range data {
		sum += strings.Count(v, "XMAS")
		sum += strings.Count(v, "SAMX")
	}
	for i := 0; i < len(data)-3; i++ {
		for j := 0; j < len(data[0])-3; j++ {
			if data[i][j] == 'X' && data[i+1][j+1] == 'M' && data[i+2][j+2] == 'A' && data[i+3][j+3] == 'S' {
				sum++
			}
			if data[i][j] == 'S' && data[i+1][j+1] == 'A' && data[i+2][j+2] == 'M' && data[i+3][j+3] == 'X' {
				sum++
			}
			if data[i][j+3] == 'X' && data[i+1][j+2] == 'M' && data[i+2][j+1] == 'A' && data[i+3][j] == 'S' {
				sum++
			}
			if data[i][j+3] == 'S' && data[i+1][j+2] == 'A' && data[i+2][j+1] == 'M' && data[i+3][j] == 'X' {
				sum++
			}
		}
	}

	return sum
}
