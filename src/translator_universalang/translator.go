package translator_universalang

import (
	"fmt"
	"regexp"
	"strconv"

	"../translator_anylang"
	"../translator_anylang_simplified"
)

//------------------------------------------------------------------------------

var constTranslationLanguages = []string{"en", "fr", "es", "zh", "da", "id", "ja", "ar", "ru"}

//------------------------------------------------------------------------------

func Init(debugVerbose bool) {
	if debugVerbose {
		fmt.Println("Load data...")
	}
	translator_anylang.Init(debugVerbose)
	loadIrregularVerbs() //TODO fix
	if debugVerbose {
		fmt.Println("  Loaded english irregular verbs.")
		fmt.Println("Load data done.")
		fmt.Println("")
	}

}

//------------------------------------------------------------------------------

func TranslateEnglishWordToUniversalang(word string, debugVerbose bool) string {

	if word == "" {
		return ""
	}

	//---
	// manage exceptions

	if tr, ok := constFullWordsTranslations[word]; ok {
		return tr.translate(debugVerbose)
	}

	//---
	// get translations

	if debugVerbose {
		fmt.Printf("\n")
	}

	var trByLanguage = make(map[string]string)
	for _, language := range constTranslationLanguages {
		var result = translator_anylang_simplified.TranslateEnglishWordSimplified(word, language, debugVerbose)
		// if isVowel([]rune(result)[0]) {result = " " + result} // add a space to keep consonants/vowels alignment TODO discuss
		if result != "" {
			trByLanguage[language] = result
		}
	}
	//trByLanguage["en"] = word

	//---
	// process conjugated forms

	if len(trByLanguage) <= 2 { // not found or almost not found
		if _, ok := trByLanguage["en"]; ok { // except in english, obviously
			// TODO identify what's a verb and what's a noun
			for reStr, translationSchema := range constRegexpTranslations {
				var re = regexp.MustCompile(reStr)
				var matches = re.FindStringSubmatch(word)
				if len(matches) > 0 {
					if len(translationSchema.wordToTranslate) >= 2 && translationSchema.wordToTranslate[0] == '$' {
						var translationSchemaCopy = translationSchema
						var integer, err = strconv.Atoi(translationSchema.wordToTranslate[1:])
						if err != nil {
							panic(err)
						}
						translationSchemaCopy.wordToTranslate = matches[integer]
						return translationSchemaCopy.translate(debugVerbose)
					} else {
						panic("matches")
					}
				}
			}
		}
	} else if len(trByLanguage) == 0 {
		panic("len(trByLanguage) == 0: " + word)
	}

	//---
	// mix up all languages into one

	var a = &aggregator{}
	a.prepare(trByLanguage)
	a.aggregate(trByLanguage)
	if debugVerbose {
		a.displayMap()
	}

	return a.finalize(debugVerbose)
}

//------------------------------------------------------------------------------

func finishesBy(word string, str string) bool {
	return len(word) > len(str) && word[len(word)-len(str):] == str
}

//------------------------------------------------------------------------------
