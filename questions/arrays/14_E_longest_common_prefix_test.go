package arrays

import (
	"fmt"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	strs := []string{"a"}
	longestCommonPrefix(strs)
}

func longestCommonPrefix(strs []string) string {

	index := 0
	ans := ""
	for {
		breakOuter := false
		var char byte
		if index < len(strs[0]) {
			char = strs[0][index]
		} else {
			breakOuter = true // handles case [""]
		}

		for i := 0; i < len(strs); i++ { // i = 0 handles case ["x"]
			if index >= len(strs[i]) {
				breakOuter = true
				break
			}
			if strs[i][index] != char {
				breakOuter = true
				break
			}
			if i == len(strs)-1 {
				ans = ans + string(char)
			}
			fmt.Println(ans)
		}
		if breakOuter {
			break
		}
		index++
	}
	return ans
}
