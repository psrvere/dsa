package arrays

import "strings"

// Approach 1 - 2 hash maps
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := map[byte]byte{}
	tMap := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		sval, sok := sMap[s[i]]
		tval, tok := tMap[t[i]]
		if !sok && !tok { // doesn't exist in both
			sMap[s[i]] = t[i]
			tMap[t[i]] = s[i]
		} else if sok && tok && (sval != t[i] || tval != s[i]) { // exists in both but values don't match
			return false
		} else if (sok && !tok) || (!sok && tok) { // exists in only one
			return false
		}
	}
	return true
}

// Approach 2 - first occurrence string transformation
func isIsomorphic2(s string, t string) bool {
	return transformString(s) == transformString(t)
}

func transformString(s string) string {
	hm := map[byte]int{}
	arr := []string{}
	for i := range s {
		if idx, ok := hm[s[i]]; ok {
			arr = append(arr, string(idx))
			continue
		}
		hm[s[i]] = i
		arr = append(arr, string(i))
	}
	return strings.Join(arr, "")
}
