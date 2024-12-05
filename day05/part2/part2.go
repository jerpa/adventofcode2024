package part2

import (
	c "github.com/jerpa/adventofcode2024/common"
	"slices"
	"sort"
	"strings"
)

func Part2(data []string) int {
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
			if !ok {
				sort.Slice(s, func(i, j int) bool {
					if d, ok := after[s[i]]; ok {
						return slices.Contains(d, s[j])
					}
					if d, ok := before[s[i]]; ok {
						return !slices.Contains(d, s[j])
					}
					return false
				})
				sum += c.GetInt(s[len(s)/2])
			}
		}

	}

	return sum
}
