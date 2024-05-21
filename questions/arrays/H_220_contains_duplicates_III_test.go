package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	tc := struct {
		i struct {
			nums      []int
			indexDiff int
			valueDiff int
		}
		o bool
	}{
		i: struct {
			nums      []int
			indexDiff int
			valueDiff int
		}{nums: []int{8, 7, 15, 1, 6, 1, 9, 15}, indexDiff: 2, valueDiff: 3},
		o: true,
	}

	ans := containsNearbyAlmostDuplicate(tc.i.nums, tc.i.indexDiff, tc.i.valueDiff)
	assert.Equal(t, tc.o, ans)
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	hm := map[int]int{}
	for i, num := range nums {
		for j := 0; j <= valueDiff; j++ {
			if j < num {
				num = num - j
			} else {
				num = j - num
			}
			if val, ok := hm[num]; ok && i-val <= indexDiff {
				return true
			}
		}
		hm[num] = i
	}
	return false
}
