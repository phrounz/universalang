package main

import (
	"fmt"
	"os"
)

//------------------------------------------------------------------------------

func main() {

	var debugVerbose = false
	if len(os.Args) == 2 {
		if os.Args[1] == "-v" {
			debugVerbose = true
		} else if os.Args[1] == "-h" {
			fmt.Printf("Usage: %s [-v]\n", os.Args[0])
			os.Exit(1)
		}
	}

	var words = []string{"i", "hello", "this", "be", "future", "past", "eat", "soccer", "language", "speak", "do", "tea"}
	for _, word := range words {
		var translation = translateEnglishWordToUniversalang(word, debugVerbose)
		fmt.Printf("%s: %s\n", word, translation)
	}
}

//------------------------------------------------------------------------------
