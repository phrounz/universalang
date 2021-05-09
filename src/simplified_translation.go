package main

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//------------------------------------------------------------------------------

func translateWordSimplified(word string, sourceLang string, targetLang string, debugVerbose bool) (resultSimplified string) {

	// translate
	var result string
	if constUseExamples {
		if sourceLang != "en" {
			panic(sourceLang)
		}
		result = exampleTranslations[word][targetLang]
	} else {
		result = translateQuery(word, sourceLang, targetLang)
	}

	var resultRaw = strings.ToLower(result)

	// replace characters
	var resultFinal1 = replaceChars(resultRaw, mTranslateCharsByLanguage[targetLang])
	var resultFinal2 = replaceChars(resultFinal1, mTranslateCharsFinal)

	// remove accents
	// https://twinnation.org/articles/33/remove-accents-from-characters-in-go
	var t = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	var err error
	resultSimplified, _, err = transform.String(t, resultFinal2)
	if err != nil {
		panic(err)
	}

	if debugVerbose {
		fmt.Printf("%s: %s > %s > %s > %s\n", targetLang, resultRaw, resultFinal1, resultFinal2, resultSimplified)
	}

	return
}

//------------------------------------------------------------------------------

func replaceChars(input string, mReplacements map[string]string) (output string) {
	for i := 0; i < len(input); {
		// TODO support more than 3 characters in mReplacements as keys?
		var r3 = input[i:min(len(input), i+3)]
		//fmt.Printf("# %s, ", output)
		if cTr, ok := mReplacements[r3]; ok {
			output += cTr
			i += 3
			//fmt.Printf("%s, %s :::\n", r3, cTr)
		} else {
			var r2 = input[i:min(len(input), i+2)]
			if cTr, ok := mReplacements[r2]; ok {
				output += cTr
				i += 2
				//fmt.Printf("%s, %s ::\n", r2, cTr)
			} else {
				var c = input[i : i+1]
				if cTr, ok := mReplacements[c]; ok {
					c = cTr
					//fmt.Printf("%s, %s :\n", c, cTr)
				}
				output += c
				i++
				//fmt.Printf(" %s \n", c)
			}
		}
	}
	return
}

//------------------------------------------------------------------------------

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

//------------------------------------------------------------------------------
