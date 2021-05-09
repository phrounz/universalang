package main

import (
	"fmt"
	"os"

	"./translator_universalang"
)

//------------------------------------------------------------------------------

func main() {

	// manages argument(s)
	var debugVerbose = false
	if len(os.Args) == 2 {
		if os.Args[1] == "-v" {
			debugVerbose = true
		} else if os.Args[1] == "-h" {
			fmt.Printf("Usage: %s [-v]\n", os.Args[0])
			os.Exit(1)
		}
	}

	// see translator_anylang/examples.go
	var words = []string{
		"i", "hello", "this", "be", "future", "past", "eat", "soccer",
		"language", "speak", "do", "tea", "yes", "no", "love"}

	for _, word := range words {
		var translation = translator_universalang.TranslateEnglishWordToUniversalang(word, debugVerbose)
		fmt.Printf("%s: %s\n", word, translation)
	}
}

//------------------------------------------------------------------------------
