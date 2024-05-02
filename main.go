package main

import (
	"bufio"
	"fmt"
	"longest-palindromic-substring/palindromes"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")

	s, err := reader.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(palindromes.LongestPalindromicSubstring(s))
}
