package palindromes

func LongestPalindromicSubstring(s string) string {
	indexSliceMap := map[rune][]int{}
	for i, v := range s {
		indexSliceMap[v] = append(indexSliceMap[v], i)
	}

	maxSize := 0
	longestString := ""
	for _, indexSlice := range indexSliceMap {
		for i := 0; i < len(indexSlice); i++ {
			for j := len(indexSlice) - 1; j >= i; j-- {
				if isPalindrome(s[indexSlice[i] : indexSlice[j]+1]) {
					if indexSlice[j]-indexSlice[i]+1 > maxSize {
						maxSize = indexSlice[j] - indexSlice[i] + 1
						longestString = s[indexSlice[i] : indexSlice[j]+1]
						break
					}
				}
			}
		}
	}

	return longestString
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}

	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[len(s)-i-1] {
			return false
		}

		i++
		j--
	}

	return true
}
