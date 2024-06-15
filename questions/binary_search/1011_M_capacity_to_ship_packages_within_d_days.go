package binarysearch

import (
	"fmt"
	"math"
)

func shipWithinDays(weights []int, days int) int {
	minWeight := max(weights)
	maxWeight := sum(weights)

	lo := minWeight
	hi := maxWeight

	for lo < hi {
		mid := lo + (hi-lo)/2 // this is our max weight for now
		fmt.Println(lo, mid, hi)
		if feasible(weights, mid, days) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	fmt.Println(lo, hi)
	return lo
}

func feasible(weights []int, mid int, days int) bool {
	daysNeeded := 1
	currentLoad := 0
	for i := range weights {
		if currentLoad+weights[i] <= mid {
			currentLoad += weights[i]
		} else {
			daysNeeded++
			currentLoad = weights[i]
		}
	}
	return daysNeeded <= days
}

func max(nums []int) int {
	max := math.MinInt64
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	fmt.Println("max: ", max)
	return max
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)
	return sum
}
