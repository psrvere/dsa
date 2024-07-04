package arrays

import "strings"

func largestGoodInteger(num string) string {
	if len(num) < 3 {
		return ""
	}

	largest := ""
	for i := 1; i < len(num)-1; i++ {
		if num[i-1] == num[i] && num[i] == num[i+1] {
			if largest == "" {
				largest = strings.Repeat(string(num[i]), 3)
			} else {
				if largest[0] < num[i] {
					largest = strings.Repeat(string(num[i]), 3)
				}
			}
		}
	}
	return largest
}
