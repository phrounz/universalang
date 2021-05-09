package translator_universalang

import (
	"fmt"

	"../translator_anylang_simplified"
)

//------------------------------------------------------------------------------

func TranslateEnglishWordToUniversalang(word string, debugVerbose bool) string {

	var languages = []string{"en", "fr", "es", "zh", "da", "id", "ja", "ar", "ru"}

	if debugVerbose {
		fmt.Printf("\n")
	}

	var trByLanguage = make(map[string]string)
	for _, language := range languages {
		var result = translator_anylang_simplified.TranslateWordSimplified(word, "en", language, debugVerbose)
		// if isVowel([]rune(result)[0]) {
		// 	result = " " + result // add a space to keep consonants/vowels alignment TODO discuss
		// }
		trByLanguage[language] = result
	}
	//trByLanguage["en"] = word

	var a = &aggregator{}
	a.prepare(trByLanguage)
	a.aggregate(trByLanguage)
	if debugVerbose {
		a.displayMap()
	}

	return a.finalize(debugVerbose)
}

//------------------------------------------------------------------------------
