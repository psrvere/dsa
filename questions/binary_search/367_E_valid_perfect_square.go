package binarysearch

// brute force O(n) - time limit exceeded
func isPerfectSquare1(num int) bool {
	hm := map[int]int{}
	for i := 2; i < num; {
		if num%i == 0 {
			hm[i]++
		} else {
			i++
		}
	}

	for _, value := range hm {
		if value%2 != 0 {
			return false
		}
	}

	return true
}

// binary search
func isPerfectSquare2(num int) bool {
	lo := 1
	hi := num
	for lo <= hi {
		mid := lo + (hi-lo)/2

		if mid*mid > num {
			hi = mid - 1
		}

		if mid*mid < num {
			lo = mid + 1
		}

		if mid*mid == num {
			return true
		}
	}
	return false
}
