package main

import (
	"fmt"
)

func main() {
	ransomNote := "bg"
	magazine := "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"
	fmt.Println(canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	// ransomNoteCharCount how many times each character appears in the ransomeNote
	ransomNoteCharCount := make([]int, 26)
	for _, r := range ransomNote {
		ransomNoteCharCount[r-'a']++
	}

	// count how many times each character appears in the magazine
	magazineCharCount := make([]int, 26)
	for _, m := range magazine {
		magazineCharCount[m-'a']++
	}

	// check if the number of times each character appears in the magazine is equal or greater than that in the ransomNote
	for i := 0; i < 26; i++ {
		if magazineCharCount[i] < ransomNoteCharCount[i] {
			return false
		}
	}

	return true
}
