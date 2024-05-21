package arrays

import (
	"sort"
	"strings"
)

// Solution 1 - group by sorted anagram string

func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}

	hm := map[string][]string{}
	for _, str := range strs {
		sorted := sortString(str)
		hm[sorted] = append(hm[sorted], str)
	}

	for _, value := range hm {
		ans = append(ans, value)
	}

	return ans
}

func sortString(str string) string {
	slice := strings.Split(str, "")
	sort.Strings(slice)
	return strings.Join(slice, "")
}

// Solution 2 - categorize by count
