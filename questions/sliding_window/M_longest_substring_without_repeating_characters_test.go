package slidingwindow

import "slices"

func lengthOfLongestSubstring(s string) int {
	// ss := []rune{}
	// for _, char := range s {
	// 	if slices.Index(ss, char) == -1 {
	// 		ss = append(ss, char)
	// 	}
	// }
	// return len(ss)
	ss := []byte{}
	max := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if slices.Index(ss, s[j]) == -1 {
				ss = append(ss, s[j])
			} else {
				break
			}
		}
		if max < len(ss) {
			max = len(ss)
		}
		ss = []byte{}
	}
	return max
}
