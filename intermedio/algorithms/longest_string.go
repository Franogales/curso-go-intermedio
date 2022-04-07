package main

import (
	"fmt"
	"strings"
)

// given an string, gets the biggest characters before it gets repeated

func main() {
	word := "abcvcpow"
	longest := getLongestString(word)
	fmt.Println(longest)
}

// here I use a function that use afor loop inside a for loop. I go I can iterate and redice it
func getLongestString(word string) (response string) {
	var tmp string
	for i := 0; i < len(word); i++ {
		letter := word[i:]
		if !strings.Contains(tmp, letter) {
			tmp += letter
			if len(tmp) > len(response) {
				response = tmp
			}
			continue
		}
		if tmp[len(tmp)-1:] != letter {
			tmp = tmp[len(tmp)-1:] + letter
			continue
		}
		tmp = letter

	}
	return response
}
