package binarySearch

import "math"

func FindIndex(inputSlice []int, value int) int {
	start := 0
	end := len(inputSlice) - 1
	var mid int
	for start <= end {
		mid = int(math.Floor((float64(start) + float64(end))/2))
		if inputSlice[mid] == value{
			return mid
		}
		if inputSlice[mid] > value {
			end = mid - 1
		} else{
			start = mid + 1
		}
	}

	return -1
}