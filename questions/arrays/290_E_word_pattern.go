package arrays

import (
	"fmt"
	"strings"
)

// Approach 1 - two hash maps
func wordPattern(pattern string, s string) bool {
	sArr := strings.Split(s, " ")

	if len(sArr) != len(pattern) {
		return false
	}

	pMap := map[int]string{}
	sMap := map[string]int{}
	for i := range pattern {
		pval, pok := pMap[int(pattern[i]-'a')]
		sval, sok := sMap[sArr[i]]

		if pok && pval == sArr[i] && sok && sval == int(pattern[i]-'a') {
			continue
		} else if (pok && pval != sArr[i]) || (sok && sval != int(pattern[i]-'a')) {
			return false
		} else if !pok && !sok {
			pMap[int(pattern[i]-'a')] = sArr[i]
			sMap[sArr[i]] = int(pattern[i] - 'a')
		}
	}
	return true
}

// Aprroach 2 - use one hash map
func wordPattern2(pattern string, s string) bool {
	hashmap := map[string]int{}
	arr := strings.Split(s, " ")
	if len(pattern) != len(arr) {
		return false
	}
	for i := range arr {
		fmt.Println(i)
		pChar := "p_" + string(pattern[i])
		sChar := "s_" + arr[i]
		if _, ok := hashmap[pChar]; !ok {
			hashmap[pChar] = i
		}

		if _, ok := hashmap[sChar]; !ok {
			hashmap[sChar] = i
		}

		if hashmap[pChar] != hashmap[sChar] {
			return false
		}
		fmt.Println(hashmap)
	}
	return true
}
