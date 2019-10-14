// Problem:
// Given a string, reverse it word by word.
//
// Example:
// Input: "coding interview does not have to have so hard"
// Output: "hard so have to have not does interview coding"
//
// Solution:
// Approach with a two-pass solution.
// The first pass is to split the string into an array of words separated by
// spaces.
// The second pass is to reverse the order of words in the array by using
// two-pointer approach: swap two values on both ends as we move toward the
// middle.
// Concatenate the values of ordered array to create a final string.
//
// Cost:
// O(n) time, O(n) space.

package leetcode

import (
	"strings"
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestReverseWordsString(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"", ""},
		{"a", "a"},
		{"a b", "b a"},
		{"a b c", "c b a"},
		{"a b c d", "d c b a"},
		{"coding interview does not have to be so hard", "hard so be to have not does interview coding"},
	}

	for _, tt := range tests {
		result := reverseWordsString(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func reverseWordsString(in string) string {
	// slice input string into a list of words separated by spaces.
	words := strings.Split(in, " ")

	// start reversing the order of words by first initializing the start and
	// end index pointer.
	start := 0
	end := len(words) - 1

	for start < end {
		// swap 2 character using a temp variable.
		tmp := words[start]
		words[start] = words[end]
		words[end] = tmp

		// move the cursor toward the middle.
		start++
		end--
	}

	// return the concatenated string created from words.
	return strings.Join(words, " ")
}
