package utils

import "strconv"

func SumStringArray(arr []string) int64 {
	var sum int64
	for _, num := range arr {
		n, _ := strconv.ParseInt(num, 10, 64)
		sum += n
	}

	return sum
}

func SumInt64Array(arr []int64) int64 {
	var sum int64
	for _, n := range arr {
		sum += n
	}

	return sum
}

func ArrayMax(arr []int64) int64 {
	var max int64
	for i := 0; i < len(arr); i++ {
		max = MaxInt(max, arr[i])
	}
	return max
}

func MaxInt(n1 int64, n2 int64) int64 {
	if n1 < n2 {
		return n2
	} else {
		return n1
	}
}
