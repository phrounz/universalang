package main

import (
	"fmt"
)

//------------------------------------------------------------------------------

func main() {

	// if len(os.Args) != 2 {
	// 	fmt.Printf("Usage: %s <google-api-token>\n", os.Args[0])
	// 	os.Exit(1)
	// }
	//
	// var token = os.Args[1]

	var words = []string{"i", "hello", "this", "be", "future", "past", "eat", "soccer", "language", "speak", "do", "tea"}
	for _, word := range words {
		var translation = translateEnglishWordToUniversalang(word)
		fmt.Printf("%s: %s\n", word, translation)
	}
}

//------------------------------------------------------------------------------
