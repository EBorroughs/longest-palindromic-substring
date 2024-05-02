package palindromes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindromeEmptyString(t *testing.T) {
	require.False(t, isPalindrome(""))
}

func TestIsPalindromeOneCharacter(t *testing.T) {
	require.True(t, isPalindrome("a"))
}

func TestIsPalindromeTrue(t *testing.T) {
	require.True(t, isPalindrome("gohangasalamiimalasagnahog"))
}

func TestIsPalindromeFalse(t *testing.T) {
	require.False(t, isPalindrome("this is not a palindrome"))
}

func TestLongestPalindromicSubstring(t *testing.T) {
	testCases := []struct {
		name           string
		value          string
		expectedString string
	}{
		{
			name:           "empty string",
			value:          "",
			expectedString: "",
		},
		{
			name:           "one character",
			value:          "s",
			expectedString: "s",
		},
		{
			name:           "palindrome",
			value:          "aracecarb",
			expectedString: "racecar",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expectedString, LongestPalindromicSubstring(tc.value))
		})
	}
}

func TestLongestPalindromicSubstringUniqueCharacters(t *testing.T) {
	require.Len(t, LongestPalindromicSubstring("abcdefghijklmnopqrstuvwxyz"), 1)
}
