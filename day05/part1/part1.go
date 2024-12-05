package part1

import (
	c "github.com/jerpa/adventofcode2024/common"
	"slices"
	"strings"
)

func Part1(data []string) int {
	sum := 0
	before := map[string][]string{}
	after := map[string][]string{}
	work := false
	for _, v := range data {
		if v == "" {
			work = true
			continue
		}
		if !work {
			s := strings.Split(v, "|")
			if _, ok := after[s[0]]; !ok {
				after[s[0]] = []string{}
			}
			after[s[0]] = append(after[s[0]], s[1])
			if _, ok := before[s[1]]; !ok {
				before[s[1]] = []string{}
			}
			before[s[1]] = append(before[s[1]], s[0])
		} else {
			s := strings.Split(v, ",")
			ok := true

			for i, cc := range s[:len(s)-1] {
				for _, vv := range s[i+1:] {
					if !slices.Contains(after[cc], vv) {
						ok = false
						break
					}
				}

				if !ok {
					break
				}
			}
			if ok {
				sum += c.GetInt(s[len(s)/2])
			}
		}

	}

	return sum
}
