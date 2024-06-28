package arrays

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		idx := abs2(nums[i])
		if nums[idx-1] > 0 {
			nums[idx-1] = -nums[idx-1]
		}
	}
	fmt.Println(nums)

	ans := []int{}
	for i := range nums {
		if nums[i] > 0 {
			ans = append(ans, i+1)
		}
	}

	return ans
}

func abs2(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
