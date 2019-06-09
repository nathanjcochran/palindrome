package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing required argument")
	}
	word := os.Args[1]

	palindromes := longestPalindromes(word)
	fmt.Println(palindromes)
}

func longestPalindromes(word string) []string {
	var (
		lengths   = make([]int, len(word)*2+1) // Keeps track of palindrome length at each position in word
		reference int                          // Index of reference palindrome in lengths slice
		maxLength int                          // Length of longest palindrome found so far
		results   []string                     // Max length palindromes found so far
	)
	for position := range lengths {
		var (
			// Word indexes for expanding out
			left  int
			right int
		)
		if position < reference+lengths[reference] {
			// If we're inside the reference palindrome, check the mirror palindrome
			mirror := reference - (position - reference)
			referenceRight := (reference-1)/2 + (lengths[reference] / 2)
			right = (position-1)/2 + (lengths[mirror] / 2)
			if referenceRight > right {
				// Mirror palindrome is completely within reference palindrome
				lengths[position] = lengths[mirror]
				continue
			} else if referenceRight < right {
				// Mirror palindrome extends beyond reference palindrome
				lengths[position] = lengths[mirror] - ((right - referenceRight) * 2)
				continue
			} else {
				// Mirror palindrome extends exactly to end of reference palindrome,
				// expand out starting at end of reference palindrome
				left = position/2 - (lengths[mirror] / 2)
			}
		} else {
			// If we're outside the reference palindrome, just start expanding
			// from either side of the position position
			left = position/2 - 1
			right = (position + 1) / 2
		}

		// Expand out
		for left >= 0 && right < len(word) && word[left] == word[right] {
			left--
			right++
		}
		lengths[position] = (right - left) - 1
		reference = position

		// Keep track of max length palindromes found so far
		if lengths[position] == maxLength {
			results = append(results, word[left+1:right])
		} else if lengths[position] > maxLength {
			maxLength = lengths[position]
			results = []string{word[left+1 : right]}
		}
	}

	return results
}
