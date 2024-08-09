package tasktwo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordFrequencyCount(t *testing.T) {
	word := "abc abc cde"
	expected := map[string]int{
		"abc": 2,
		"cde": 1,
	}
	actualResult := wordFrequencyCount(word)
	assert.Equal(t, expected, actualResult, expected)
}

func TestCheckPalindrome(t *testing.T) {
	palindromeWord := "abccba"
	notPalindromeWord := "abcab"
	assert.Equal(t, CheckPalindrome(palindromeWord), true)
	assert.Equal(t, CheckPalindrome(notPalindromeWord), false)

}
