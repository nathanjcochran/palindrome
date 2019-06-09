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
		lengths = make([]int, len(word)*2+1)
		ref     int
		max     int
	)

	for i := range lengths {
		if i >= ref+lengths[ref] {
			// If we're outside the reference palindrome, expand
			var (
				left  = i/2 - 1
				right = (i + 1) / 2
			)
			for left >= 0 && right < len(word) && word[left] == word[right] {
				left--
				right++
			}
			lengths[i] = (right - left) - 1
			ref = i
			if lengths[i] > max {
				max = lengths[i]
			}
		} else {
			// Otherwise, check mirror palindrome
			mir := ref - (i - ref)
			refRight := (ref-1)/2 + (lengths[ref] / 2)
			right := (i-1)/2 + (lengths[mir] / 2)
			if refRight > right {
				// Mirror palindrome is completely within reference palindrome
				lengths[i] = lengths[mir]
			} else if refRight < right {
				// Mirror palindrome extends beyond reference palindrome
				lengths[i] = lengths[mir] - ((right - refRight) * 2)
			} else {
				// Mirror palindrome extends exactly to end of reference palindrome
				left := i/2 - (lengths[mir] / 2)
				for left >= 0 && right < len(word) && word[left] == word[right] {
					left--
					right++
				}
				lengths[i] = (right - left) - 1
				ref = i
				if lengths[i] > max {
					max = lengths[i]
				}
			}
		}
	}

	var results []string
	for i, length := range lengths {
		if length == max {
			left := i/2 - length/2
			right := (i+1)/2 + length/2
			results = append(results, word[left:right])
		}
	}
	return results
}
