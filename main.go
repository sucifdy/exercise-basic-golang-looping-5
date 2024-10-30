package main

import (
	"fmt"
	"strings"
	"unicode"
)

// ReverseWord function that reverses each word in a string
func ReverseWord(str string) string {
	// Split string into words
	words := strings.Fields(str)

	// Iterate over each word and reverse it
	for i, word := range words {
		// Reverse the word
		reversedWord := reverseString(word)

		// Adjust capitalization
		if isAllUpper(word) {
			// If the original word is all uppercase, keep it all uppercase
			reversedWord = strings.ToUpper(reversedWord)
		} else if unicode.IsUpper(rune(word[0])) {
			// If original word starts with uppercase, adjust first letter in reversed word
			reversedWord = strings.ToUpper(string(reversedWord[0])) + strings.ToLower(reversedWord[1:])
		}

		// Update the word in the slice
		words[i] = reversedWord
	}

	// Join words back into a single string with spaces
	return strings.Join(words, " ")
}

// reverseString helper function to reverse a string
func reverseString(s string) string {
	// Convert string to rune slice to handle Unicode properly
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// isAllUpper checks if all characters in a string are uppercase
func isAllUpper(word string) bool {
	for _, r := range word {
		if unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(ReverseWord("Aku Sayang Ibu"))        // Output: "Uka Gnayas Ubi"
	fmt.Println(ReverseWord("ini terlalu mudah"))     // Output: "ini ulalret hadum"
	fmt.Println(ReverseWord("KITA SELALU BERSAMA"))   // Output: "ATIK ULALES AMASREB"
}
