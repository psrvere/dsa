package arrays

// Approach 1 - brute force

func destCity1(paths [][]string) string {
	for i := 0; i < len(paths); i++ {
		destination := paths[i][1]
		final := true
		for j := 0; j < len(paths); j++ {
			source := paths[j][0]
			if destination == source {
				final = false
				break
			}
		}
		if final {
			return destination
		}
	}
	return ""
}

// Approach 2 - hashmap
// TC: O(n)

func destCity2(paths [][]string) string {
	sourceMap := map[string]bool{}
	for _, path := range paths {
		sourceMap[path[0]] = true
	}

	for _, path := range paths {
		destination := path[1]
		if _, ok := sourceMap[destination]; !ok {
			return destination
		}
	}

	return ""
}
