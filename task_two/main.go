package tasktwo

import (
	"strings"
	"unicode"
)

func getWord(s string) []string {
	words := strings.FieldsFunc(strings.ToLower(s), func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})

	return words

}

func removePunction(s string) string {
	cur_word := ""
	for _, word := range s {
		if unicode.IsLetter(word) || unicode.IsNumber(word) {
			cur_word += strings.ToLower(string(word))
		}

	}

	return cur_word
}

func wordFrequencyCount(s string) map[string]int {

	counter := make(map[string]int)
	words := getWord(s)

	for _, word := range words {
		counter[word] += 1
	}

	return counter
}

func CheckPalindrome(s string) bool {

	cur_word := removePunction(s)

	left, right := 0, len(s)-1

	for left <= right {
		if cur_word[left] != cur_word[right] {
			return false
		}

		left += 1
		right -= 1
	}

	return true
}
