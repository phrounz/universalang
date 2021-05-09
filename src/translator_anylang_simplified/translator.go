package translator_anylang_simplified

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"../translator_anylang"
)

//------------------------------------------------------------------------------

func TranslateEnglishWordSimplified(word string, targetLang string, debugVerbose bool) (resultSimplified string) {

	// translate
	var result, isAlreadyIPA = translator_anylang.Translate(word, "en", targetLang, debugVerbose)

	// remove uppercase
	var resultRaw = strings.ToLower(result)

	// replace characters (result => IPA)
	var result1 string
	if isAlreadyIPA {
		result1 = resultRaw
	} else {
		result1 = replaceChars(resultRaw, mTranslateCharsByLanguage[targetLang])
	}

	// replace characters (IPA => readable english alphabet (with accents though))
	var result2 = replaceChars(result1, mTranslateCharsFinal)

	// remove accents
	// https://twinnation.org/articles/33/remove-accents-from-characters-in-go
	var t = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	var err error
	resultSimplified, _, err = transform.String(t, result2)
	if err != nil {
		panic(err)
	}

	if debugVerbose {
		fmt.Printf("%s: %s > %s > %s > %s\n", targetLang, resultRaw, result1, result2, resultSimplified)
	}

	return
}

//------------------------------------------------------------------------------

// TODO support more than 3 characters in mReplacements as keys?
func replaceChars(input string, mReplacements map[string]string) (output string) {
	for i := 0; i < len(input); {
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
