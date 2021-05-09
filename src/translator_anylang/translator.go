package translator_anylang

//------------------------------------------------------------------------------

const constUseExamples = true

//------------------------------------------------------------------------------

// isIPA: true if the result is already in International Phonetic Alphabet
func Translate(word string, sourceLang string, targetLang string, debugVerbose bool) (result string, isIPA bool) {
	if constUseExamples {
		if sourceLang != "en" {
			panic(sourceLang)
		}
		var ok bool
		if result, ok = exampleTranslations[word][targetLang]; !ok {
			panic("missing translation of " + word + " in " + targetLang + " in examples")
		}
		isIPA = (targetLang == "en")
	} else {
		result = translateQueryGoogle(word, sourceLang, targetLang)
		isIPA = false // TODO
	}
	return
}

//------------------------------------------------------------------------------
