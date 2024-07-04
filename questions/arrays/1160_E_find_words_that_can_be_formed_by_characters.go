package arrays

import "fmt"

// case 1 - character in word is not in chars
// case 2 - character count in word is more than character count in chars

// Approach 1
func countCharacters(words []string, chars string) int {
	cMap := map[rune]int{}
	for _, char := range chars {
		cMap[char]++
	}

	count := 0
	for _, word := range words {
		mapCopy := map[rune]int{}
		copyMap(cMap, mapCopy)

		skip := false
		for _, w := range word {
			if _, ok := mapCopy[w]; !ok {
				skip = true
				break // case 1
			} else {
				mapCopy[w]--
			}
		}
		fmt.Println("skip: ", skip)
		if skip {
			continue // case 1
		}

		for _, value := range mapCopy {
			if value < 0 { // case 2
				skip = true
				break
			}
		}
		fmt.Println("skip: ", skip)
		if skip {
			continue // case 2
		}

		// valid cases
		count += len(word)
	}

	return count
}

func copyMap(a, b map[rune]int) {
	for key, value := range a {
		b[key] = value
	}
}

// Approach 2 - shorter code
// case 1 - character in word is not in chars
// case 2 - character count in word is more than character count in chars
func countCharacters2(words []string, chars string) int {
	cMap := map[rune]int{}
	for _, char := range chars {
		cMap[char]++
	}

	count := 0
	for _, word := range words {
		wMap := map[rune]int{}
		for _, w := range word {
			wMap[w]++
		}

		skip := false
		for key, val := range wMap {
			if val > cMap[key] {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		// valid cases
		count += len(word)
	}
	return count
}
