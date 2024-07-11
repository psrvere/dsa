package arrays

// Approach 1 - self
func maxScore(s string) int {
	leftSum := 0
	rightSum := 0
	if string(s[0]) == "0" {
		leftSum++
	}
	for i := 1; i < len(s); i++ {
		if string(s[i]) == "1" {
			rightSum++
		}
	}
	maxSum := leftSum + rightSum

	for i := 1; i < len(s)-1; i++ {
		if string(s[i]) == "0" {
			leftSum++
		} else {
			rightSum--
		}
		sum := leftSum + rightSum
		if maxSum < sum {
			maxSum = sum
		}
	}
	return maxSum
}
