package arrays

import "fmt"

// Approach 1 - brute force O(n2)
func numIdenticalPairs(nums []int) int {
	count := 0
	for l := 0; l < len(nums); l++ {
		for r := l + 1; r < len(nums); r++ {
			if nums[l] == nums[r] {
				count++
			}
		}
	}
	return count
}

// Approach 2 -hashmap and combinatorics
func numIdenticalPairs2(nums []int) int {
	hm := map[int]int{}
	for _, num := range nums {
		hm[num]++
	}
	count := 0
	for _, val := range hm {
		// count of pairs is nC2
		if val >= 2 {
			count += getnC2(val)
		}
	}
	return count
}

func getnC2(n int) int {
	return n * (n - 1) / 2
}

// Approach 3 - Only hashmap
func numIdenticalPairs3(nums []int) int {
	hm := map[int]int{}
	count := 0
	for _, num := range nums {
		val := hm[num]
		fmt.Println(val)
		if val < 1 {
			hm[num]++
			continue
		}
		count += val
		hm[num]++
	}
	return count
}
