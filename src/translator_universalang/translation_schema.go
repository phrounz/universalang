package translator_universalang

//------------------------------------------------------------------------------

const constPastTenseSuffix = "*"
const constPastParticipleSuffix = constPastTenseSuffix
const constToBePresentForm = "ei"

//------------------------------------------------------------------------------

// either fullTranslation must be provided
// or wordToTranslate (optionally with prefix and suffix)
type translationSchema struct {
	fullTranslation string
	wordToTranslate string
	prefix          string
	suffix          string
}

func (t *translationSchema) translate(debugVerbose bool) string {
	if t.fullTranslation != "" {
		return t.fullTranslation
	}
	return t.prefix + TranslateEnglishWordToUniversalang(t.wordToTranslate, debugVerbose) + t.suffix
}

var constFullWordsTranslations = map[string]translationSchema{
	"a":    translationSchema{fullTranslation: ""},
	"an":   translationSchema{fullTranslation: ""},
	"the":  translationSchema{fullTranslation: ""},
	"am":   translationSchema{fullTranslation: constToBePresentForm},
	"is":   translationSchema{fullTranslation: constToBePresentForm},
	"are":  translationSchema{fullTranslation: constToBePresentForm},
	"was":  translationSchema{fullTranslation: constToBePresentForm + constPastTenseSuffix},
	"were": translationSchema{fullTranslation: constToBePresentForm + constPastTenseSuffix},
}

var constRegexpTranslations = map[string]translationSchema{
	"^(.*)ing$":     translationSchema{wordToTranslate: "$1", suffix: "~"},                  // gerund
	"^(.*)[^\\']s$": translationSchema{wordToTranslate: "$1", suffix: "$"},                  // plural
	"^(.*)ed$":      translationSchema{wordToTranslate: "$1", suffix: constPastTenseSuffix}, // past tense
	"^(.*)\\'s$":    translationSchema{wordToTranslate: "$1", suffix: " ed"},                // possessive
}

//------------------------------------------------------------------------------
