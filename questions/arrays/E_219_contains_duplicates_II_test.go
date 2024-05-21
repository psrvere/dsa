package arrays

import "slices"

// Solution 1 - two loops - O(n2)

func containsNearbyDuplicate1(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && j-i <= k && i != j {
				return true
			}
		}
	}
	return false
}

// Solution 2 - maintain a slice of k length

func containsNearbyDuplicate2(nums []int, k int) bool {
	if k == 0 {
		return true
	}
	arr := []int{}
	for _, num := range nums {
		if slices.Index(arr, num) == -1 {
			if len(arr) < k {
				arr = append(arr, num)
			} else {
				arr = arr[1:] // remove first element
				arr = append(arr, num)
			}
		} else {
			return true
		}
	}
	return false
}

// Solution 3 - use an hash map

func containsNearbyDuplicate3(nums []int, k int) bool {
	m := map[int]int{}
	for i, num := range nums {
		if v, ok := m[num]; ok && i-v <= k {
			return true
		}
		m[num] = i
	}
	return false
}
