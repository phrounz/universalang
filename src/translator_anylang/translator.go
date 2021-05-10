package translator_anylang

//------------------------------------------------------------------------------

type EnumMode int

const (

	// In this mode, only these words are available: (see examples.go)
	// "hello", "i", "you", "this", "be", "future", "past", "eat", "soccer", "language", "speak", "do", "tea", "yes", "no", "love"
	enumModeExamples EnumMode = iota

	// loads dictionary files from ./data/
	enumModeOfflineDictionaries

	// Use Google Translate. It's not tested and requires a Google Translate API token.
	enumModeGoogleTranslate
)

var globalEnumMode = enumModeOfflineDictionaries

//------------------------------------------------------------------------------

// isIPA: true if the result is already in International Phonetic Alphabet
func Translate(word string, sourceLang string, targetLang string, debugVerbose bool) (result string, isIPA bool) {
	if globalEnumMode == enumModeExamples {
		if sourceLang != "en" {
			panic(sourceLang)
		}
		var ok bool
		if result, ok = exampleTranslations[word][targetLang]; !ok {
			panic("missing translation of " + word + " in " + targetLang + " in examples")
		}
		isIPA = (targetLang == "en")

	} else if globalEnumMode == enumModeOfflineDictionaries {
		if sourceLang != "en" {
			panic(sourceLang)
		}
		if targetLang == "en" {
			result = word
			isIPA = false
		} else {
			result = translateFromEnglishUsingOfflineDictionary(word, targetLang)
			isIPA = false
		}

	} else if globalEnumMode == enumModeGoogleTranslate {
		result = translateQueryGoogle(word, sourceLang, targetLang)
		isIPA = false // TODO

	} else {
		panic(globalEnumMode)
	}
	return
}

//------------------------------------------------------------------------------
