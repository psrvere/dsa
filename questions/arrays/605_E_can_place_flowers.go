package arrays

// Approach 1 - self
func canPlaceFlowers(flowerbed []int, n int) bool {
	// identify patterns of 0
	// first, atleast 000 in the middle - can plant 1
	// second, atleast 00 at the end - can plant 1

	// corner case
	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		return true
	}
	streak := 0
	planted := 0
	for i := 0; i < len(flowerbed); i++ {
		// update streak
		if flowerbed[i] == 0 {
			streak++
		} else {
			streak = 0
		}

		// corner cases
		if (i == 1 || i == len(flowerbed)-1) && streak == 2 {
			planted++
			streak = 1
		}

		if streak == 3 {
			planted++
			streak = 1
		}
	}
	return n <= planted
}
