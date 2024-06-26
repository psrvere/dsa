package arrays

import "fmt"

// two pointers - Did not work
func pivotIndex(nums []int) int {
	l, r := 0, len(nums)-1
	lsum, rsum := getLeftValue(nums, l), getRightValue(nums, r)
	for l <= r {
		fmt.Println("l,r: ", l, r)
		if lsum < rsum {
			fmt.Println("l++")
			l++
			lsum += getLeftValue(nums, l)
		} else if lsum > rsum {
			fmt.Println("r--")
			r--
			rsum += getRightValue(nums, r)
		} else if lsum == rsum && l == r {
			return l
		} else if lsum == rsum && l < r {
			fmt.Println("l++")
			l++
			lsum += getLeftValue(nums, l)
		}
		fmt.Println("sum: ", lsum, rsum)
	}
	if lsum == rsum {
		return l
	} else {
		return -1
	}
}

func getLeftValue(nums []int, i int) int {
	if i == 0 {
		return 0
	}
	return nums[i-1]
}

func getRightValue(nums []int, i int) int {
	if i == len(nums)-1 {
		return 0
	}
	return nums[i+1]
}

// single iteration
func pivotIndex2(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	lsum, rsum := 0, 0
	for i := 0; i < len(nums); i++ {
		lsum += getLeftValue(nums, i)
		rsum = sum - lsum - nums[i]
		if lsum == rsum {
			return i
		}
	}
	return -1
}
