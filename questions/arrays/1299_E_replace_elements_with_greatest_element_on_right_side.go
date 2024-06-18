package arrays

// Approach 1 - iterate from beginning
func replaceElements1(arr []int) []int {
	for i := range arr {
		if i == len(arr)-1 {
			arr[i] = -1
			break
		}
		arr[i] = max(arr[i+1:])
	}
	return arr
}

func max(arr []int) int {
	max := 0
	for i := range arr {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

// Approach 2 - iterate from end
func replaceElements2(arr []int) []int {
	n := len(arr)
	maxFromLast := arr[n-1]
	arr[n-1] = -1
	for i := n - 2; i >= 0; i-- {
		value := arr[i]
		arr[i] = maxFromLast
		if maxFromLast < value {
			maxFromLast = value
		}
	}
	return arr
}

// Approach 21 - slightly better
func replaceElements(arr []int) []int {
	n := len(arr)
	maxFromLast := -1
	for i := n - 1; i >= 0; i-- {
		value := arr[i]
		arr[i] = maxFromLast
		if maxFromLast < value {
			maxFromLast = value
		}
	}
	return arr
}
