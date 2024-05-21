package arrays

// Solution1 - frequency counter: increase map count for all chars in s, decrease map count for all chars in t - O(n)

func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		sMap[s[i]] = sMap[s[i]] + 1
		sMap[t[i]] = sMap[t[i]] - 1
	}

	for _, val := range sMap {
		if val != 0 {
			return false
		}
	}
	return true
}

// Solution 2 - sorting - O(nlogn)

// Follow up - what if character contained unicode character?
