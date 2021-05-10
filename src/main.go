package main

import (
	"fmt"
	"os"
	"strings"

	"./translator_universalang"
)

//------------------------------------------------------------------------------

func main() {

	// manages argument(s)
	var debugVerbose = false
	var sentence = ""

	if len(os.Args) >= 2 {
		for i, arg := range os.Args {
			if i > 0 {
				if arg == "-v" {
					debugVerbose = true

				} else if arg == "-h" {
					usage()

				} else {
					sentence += arg + " "
				}
			}
		}
	}

	if sentence == "" {
		usage()
	}

	translator_universalang.Init(debugVerbose)

	var words = strings.Split(sentence, " ")
	var translationStr = ""
	for _, word := range words {
		word = strings.ToLower(word)
		var addedPunctuation = ""
		if len(word) > 1 {
			var c = word[len(word)-1]
			if c == ',' || c == '.' || c == '?' || c == '!' || c == ';' {
				word = word[:len(word)-1]
				addedPunctuation = string(c)
			}
		}
		if word != "" {
			var translation = translator_universalang.TranslateEnglishWordToUniversalang(word, debugVerbose)
			if translation != "" {
				if translationStr != "" {
					translationStr += " "
				}
				translationStr += translation
			}
			if addedPunctuation != "" {
				translationStr += addedPunctuation
			}
		}

	}
	fmt.Printf("%s\n", translationStr)
}

//------------------------------------------------------------------------------

func usage() {
	fmt.Printf("Usage: %s [-v] [sentence]\n", os.Args[0])
	os.Exit(1)
}
