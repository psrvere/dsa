package binarysearch

import "fmt"

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		fmt.Println(low, high, mid)
		if nums[mid] > target {
			high = mid - 1
		}

		if nums[mid] < target {
			low = mid + 1
		}

		if nums[mid] == target {
			return mid
		}
	}
	return -1
}
