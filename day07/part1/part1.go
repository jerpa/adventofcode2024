package part1

import (
	c "github.com/jerpa/adventofcode2024/common"
	"strings"
)

func Part1(data []string) int {
	sum := 0
	for _, v := range data {
		s := strings.Split(v, ": ")
		target := c.GetInt(s[0])
		nums := c.GetInts(strings.Split(s[1], " "))
		if calc(nums, 0, target) {
			sum += target
		}

	}

	return sum
}
func calc(nums []int, sum, target int) bool {
	if len(nums) == 0 {
		return sum == target
	}
	if calc(nums[1:], sum+nums[0], target) {
		return true
	}
	if calc(nums[1:], sum*nums[0], target) {
		return true
	}
	return false
}
