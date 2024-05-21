package arrays

// https://leetcode.com/problems/contains-duplicate/description/

// Solution 1 - compare every two elements - two loops - O(n2)
// Solution 2 - sort the array and compare every adjacent element pair - one loop - O(nlogn)
// Solution 3 - use hash map - O(n)

func containsDuplicate(nums []int) bool {
	m := map[int]bool{}
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = true
	}
	return false
}
