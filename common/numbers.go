package common

import "strconv"

// GetInts converts a slice of strings to a slice of ints
func GetInts(data []string) []int {
	result := []int{}
	for _, v := range data {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, val)
	}
	return result
}

// GetInt get the int value of a string
func GetInt(data string) int {
	val, err := strconv.Atoi(data)
	if err != nil {
		panic(err.Error())
	}
	return val
}

// GetInt64 get the int64 value of a string
func GetInt64(data string) int64 {
	val, err := strconv.Atoi(data)
	if err != nil {
		panic(err.Error())
	}
	return int64(val)
}

// GetFloats converts a slice of strings to a slice of floats
func GetFloats(data []string) []float64 {
	result := []float64{}
	for _, v := range data {
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, val)
	}
	return result
}

// MinInt returns the lowest number from a slice
func MinInt(nums []int) int {
	res := 0
	for i, v := range nums {
		if i == 0 || v < res {
			res = v
		}
	}
	return res
}

// MaxInt returns the highest number from a slice
func MaxInt(nums []int) int {
	res := 0
	for i, v := range nums {
		if i == 0 || v > res {
			res = v
		}
	}
	return res
}

// Sum returns the sum of a slice of ints
func Sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// AbsInt returns the absolute values for two integers as an integer
func AbsInt(num1, num2 int) int {
	num1 = num1 - num2
	if num1 < 0 {
		num1 *= -1
	}
	return num1
}

// ConsecutiveSum returns the sum of all the previous values from 1..num
func ConsecutiveSum(num int) int {
	if num == 0 {
		return 0
	}
	v := float32(num)
	return int(v * ((1 + v) / 2))
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(x ...int) int {
	if len(x) == 1 {
		return x[0]
	} else if len(x) > 2 {
		return LCM(x[0], LCM(x[1:]...))
	}
	return x[0] * x[1] / GCD(x[0], x[1])
}
