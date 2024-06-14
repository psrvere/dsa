package binarysearch

func mySqrt(x int) int {
	lo := 1
	hi := x
	for lo <= hi {
		mid := lo + (hi-lo)/2
		square := mid * mid

		if square > x {
			hi = mid - 1
		}

		if square < x {
			lo = mid + 1
		}

		if square == x {
			return mid
		}
	}
	return hi
}
