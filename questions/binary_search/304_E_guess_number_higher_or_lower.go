package binarysearch

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(int) int {
	// dummy function
	return 0
}

func guessNumber(n int) int {
	lo := 1
	hi := n
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if guess(mid) == -1 {
			hi = mid - 1
		}

		if guess(mid) == 1 {
			lo = mid + 1
		}

		if guess(mid) == 0 {
			return mid
		}
	}
	return n
}
