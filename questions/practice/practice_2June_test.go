package practice

import (
	"slices"
	"strings"
)

// Arrays
// contains duplicate - 217 - easy
func containsDuplicate(nums []int) bool {
	hm := map[int]bool{}
	for _, num := range nums {
		if _, ok := hm[num]; ok {
			return true
		}
		hm[num] = true
	}
	return false
}

// contains duplicate II - 219 - easy
// store index in the hash map
func containsNearbyDuplicate_2June(nums []int, k int) bool {
	// loop through the numbers in numbs
	// maintain a hash map of size k
	// if a character is found in the array then it contains nearby duplicates
	if k == 0 {
		return false
	}

	hm := map[int]int{}
	for i, num := range nums {
		// case 1 - num not found in hm, added to hm
		// case 2 - num found in hm, i - j <= k, return true
		// case 3 - num found in hm, i - j > k, hm updated with latest index

		if j, ok := hm[num]; ok && i-j <= k {
			return true
		}
		hm[num] = i
	}

	return false
}

// contains duplicates III - 220 - hard
// in this you won't be able to find exact number from hash map because numbers are almost duplicate
// so we have to search many number in hash map

// Approach 1 - two loops - O(n2)
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	hm := map[int]int{}
	for i, num := range nums {
		for val := range indexDiff {
			diff := 0
			if val < num {
				diff = num - val
			} else {
				diff = val - num
			}

			if j, ok := hm[diff]; ok && diff <= valueDiff && i-j <= indexDiff {
				return true
			}
		}
		hm[num] = i
	}
	return false
}

// Approach 2 - Bucketing

func containsNearbyAlmostDuplicate_2(nums []int, indexDiff int, valueDiff int) bool {
	bucket := map[int]int{}
	bs := valueDiff + 1
	for i, num := range nums {
		id := num / bs

		if num < 0 {
			id--
		}

		// check in the same bucket
		if j, ok := bucket[id]; ok && i-j <= indexDiff && isValueDiffFound(num, nums[j], valueDiff) {
			return true
		}

		// check in once previous bucket
		if j, ok := bucket[id-1]; ok && i-j <= indexDiff && isValueDiffFound(num, nums[j], valueDiff) {
			return true
		}

		// check in the one next bucket
		if j, ok := bucket[id+1]; ok && i-j <= indexDiff && isValueDiffFound(num, nums[j], valueDiff) {
			return true
		}

		bucket[id] = i
	}
	return false
}

func isValueDiffFound(x int, y int, valueDiff int) bool {
	diff := 0
	if x < y {
		diff = y - x
	} else {
		diff = x - y
	}
	if diff <= valueDiff {
		return true
	}
	return false
}

// Approach 3 - Binary Search Tree - TBD

// Valid anagram - 242 - easy
func isAnagram_2June(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hm := map[byte]int{}
	for i := range s {
		hm[s[i]]++
		hm[t[i]]--
	}
	for _, val := range hm {
		if val != 0 {
			return false
		}
	}
	return true
}

// Two sum - 1 - easy
func twoSum(nums []int, target int) []int {
	hm := map[int]int{}
	for i, num := range nums {
		complement := target - num
		if j, ok := hm[complement]; ok {
			return []int{i, j}
		}
		hm[num] = i
	}
	return nil
}

// group anagrams - 49 - medium
// Approach 1 - two strings are anagram is their sorted strings are equal
// TC - O(n * mlogm), SC - O(NM)
func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	hm := map[string][]string{}
	for _, str := range strs {
		sorted := sortString(str)
		hm[sorted] = append(hm[sorted], str)
	}

	for _, val := range hm {
		ans = append(ans, val)
	}

	return ans
}

func sortString(str string) string {
	slice := strings.Split(str, "")
	slices.Sort(slice)
	return strings.Join(slice, "")
}

// Approach 2 -
