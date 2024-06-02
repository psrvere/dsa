package arrays

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
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

func Test_groupAnagrams2(t *testing.T) {
	tc := struct {
		i []string
		o [][]string
	}{
		i: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
		o: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
	}

	ans := groupAnagrams2(tc.i)
	fmt.Println("ans: ", ans)
	// assert.Equal(t, true, reflect.DeepEqual(ans, tc.o))
}

func groupAnagrams2(strs []string) [][]string {
	ans := map[string][]string{}
	count := make([]int, 26)
	for _, s := range strs {
		for i := 0; i < 26; i++ {
			count[i] = 0
		}
		for _, c := range s {
			count[c-'a']++
		}
		key := ""
		for i := 0; i < 26; i++ {
			key += "#"
			key += strconv.Itoa(count[i])
		}
		if _, ok := ans[key]; !ok {
			ans[key] = []string{}
		}
		ans[key] = append(ans[key], s)
	}
	result := [][]string{}
	for _, val := range ans {
		result = append(result, val)
	}
	return result
}

// Solution 2 - categorize by count - another version

func Test_groupAnagrams2a(t *testing.T) {
	tc := struct {
		i []string
		o [][]string
	}{
		i: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
		o: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
	}

	ans := groupAnagrams2a(tc.i)
	fmt.Println("ans: ", ans)
	// assert.Equal(t, true, reflect.DeepEqual(ans, tc.o))
}

func hash(s string) string {
	res := make([]byte, 26)
	for _, char := range s {
		// char := s[i]
		res[char-'a']++
	}
	return string(res)
}

func groupAnagrams2a(strs []string) [][]string {
	res := [][]string{}
	m := make(map[string]int)
	for _, w := range strs {
		h := hash(w)
		idx, ok := m[h]
		if ok {
			res[idx] = append(res[idx], w)
		} else {
			res = append(res, []string{w})
			m[h] = len(res) - 1
		}
	}

	return res
}
