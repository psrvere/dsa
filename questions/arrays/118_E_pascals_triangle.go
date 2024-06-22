package arrays

// learning: when arrays are created with make
// use indexes to add element to array
// using append will increase the size of array

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	// ans = append(ans, []int{1})
	ans[0] = []int{1} // first element
	for i := 1; i < numRows; i++ {
		arr := make([]int, i+1)
		for j := 0; j < len(arr); j++ {
			if j == 0 {
				arr[j] = ans[i-1][0] // first element of above array
			} else if j == len(arr)-1 {
				arr[j] = ans[i-1][len(ans[i-1])-1] // last element of above array
			} else {
				arr[j] = ans[i-1][j-1] + ans[i-1][j]
			}
		}
		ans[i] = arr
	}
	return ans
}
