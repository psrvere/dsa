package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Solution 1 - two pointers - timelimit exeeded O(n2)
// Solution 2 - below solutions - O(n)

func Test_isPalindrome1(t *testing.T) {
	tc := struct {
		i []int
		o int
	}{
		i: []int{7, 1, 5, 3, 6, 4},
		o: 5,
	}

	max := maxProfit(tc.i)
	assert.Equal(t, tc.o, max)
}

func maxProfit(prices []int) int {
	minPrice := int(^uint(0) >> 1)
	// fmt.Println(uint(0))
	// fmt.Printf("%d %b\n", ^uint(0), ^uint(0))
	// fmt.Printf("%d %b\n", ^uint(0)>>1, ^uint(0)>>1)
	// fmt.Printf("%b\n", ^uint(0))
	// fmt.Printf("%b\n", ^uint(0)>>1)
	// fmt.Printf("%b %b %b\n", 21, 21>>1, 21<<1)
	// fmt.Printf("%d %d %d\n", 21, 21>>1, 21<<1)
	// fmt.Printf("%d %b\n", ^uint(0) << 1)
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
