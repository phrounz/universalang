package main

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//------------------------------------------------------------------------------

const constUseExamples = true

//------------------------------------------------------------------------------

var translations = map[string]map[string]string{
	"hello": map[string]string{
		"en": "həˈlō",
		"fr": "bonjour",
		"es": "hola",
		"zh": "nǐ hǎo",
		"da": "hej",
		"id": "halo",
		"ja": "kon'nichiwa",
		"ar": "marhabaan",
		"ru": "privet",
	},
	"i": map[string]string{
		"en": "i", // ?? TODO
		"fr": "je",
		"es": "yo",
		"zh": "wǒ",
		"da": "jeg",
		"id": "saya",
		"ja": "watashi",
		"ar": "'ana",
		"ru": "ya",
	},
	"be": map[string]string{
		"en": "bē",
		"fr": "être",
		"es": "sí",
		"zh": "shì",
		"da": "ja",
		"id": "iya",
		"ja": "arimasu",
		"ar": "yakun",
		"ru": "byt'",
	},
	"eat": map[string]string{
		"en": "ēt",
		"fr": "manger",
		"es": "comer",
		"zh": "chī",
		"da": "spise",
		"id": "makan",
		"ja": "taberu",
		"ar": "takul",
		"ru": "yest'",
	},
	"this": map[string]string{
		"en": "T͟His",
		"fr": "cette", //ce
		"es": "esta",  //esto
		"zh": "zhè",
		"da": "det her",
		"id": "ini",
		"ja": "kono",
		"ar": "hadha",
		"ru": "eto",
	},
	"past": map[string]string{
		"en": "past",
		"fr": "passé",
		"es": "pasado",
		"zh": "guòqù",
		"da": "fortiden",
		"id": "masa lalu",
		"ja": "kako",
		"ar": "madi",
		"ru": "byt'",
	},
	"future": map[string]string{
		"en": "ˈfyo͞oCHər",
		"fr": "futur",
		"es": "futuro",
		"zh": "wèilái",
		"da": "fremtiden",
		"id": "masa depan",
		"ja": "mirai",
		"ar": "mustaqbal",
		"ru": "",
	},
	"soccer": map[string]string{
		"en": "ˈsäkər",
		"fr": "football",
		"es": "fútbol",
		"zh": "zúqiú",
		"da": "fodbold",
		"id": "sepak bola",
		"ja": "sakkā",
		"ar": "kurat alqadam",
		"ru": "budushcheye",
	},
	"language": map[string]string{
		"en": "ˈlaNGɡwij",
		"fr": "langue",
		"es": "idioma",
		"zh": "yǔ",
		"da": "sprog",
		"id": "bahasa",
		"ja": "gengo",
		"ar": "lugha",
		"ru": "",
	},
	"speak": map[string]string{
		"en": "spēk",
		"fr": "parler",
		"es": "hablar",
		"zh": "shuōhuà",
		"da": "tale",
		"id": "berbicara",
		"ja": "hanasu",
		"ar": "tahduth",
		"ru": "govorit'",
	},
	"do": map[string]string{
		"en": "do͞o",
		"fr": "faire",
		"es": "hacer",
		"zh": "qù zuò",
		"da": "gør",
		"id": "melakukan",
		"ja": "suru",  // okonau
		"ar": "faeal", // lukaa yafeal
		"ru": "delat'",
	},
	"tea": map[string]string{
		"en": "tē",
		"fr": "thé",
		"es": "té",
		"zh": "chá",
		"da": "te",
		"id": "teh",
		"ja": "ocha",
		"ar": "shay",
		"ru": "chay",
	},
}

//------------------------------------------------------------------------------

// TODO return international phonetic alphabet as output
// instead of trying to eliminate IPA characters.
// https://en.wikipedia.org/wiki/International_Phonetic_Alphabet

// Maybe try to find translation lists e.g. https://github.com/jmbeach/duolingo-vocab-lists/

var mTranslateChars = map[string]string{
	"ä": "o",
	"ə": "e",
	"ē": "i",
	"ø": "o",
	"œ": "e",
	"ˈ": "",

	"ç": "ss",
	" ": "",
	"'": "",
}

var mTranslateCharsByLanguage = map[string]map[string]string{
	// https://en.wikipedia.org/wiki/Help:IPA/Mandarin
	"zh": map[string]string{
		"g": "k",
		"z": "ts",
		"c": "tsh", //tsʰ
		// TODO
	},
	// https://en.wikipedia.org/wiki/Help:IPA/French
	"fr": map[string]string{
		"c": "k",
		"ç": "s",
		// TODO
	},
	// TODO other languages
}

func translateWordSimplified(word string, sourceLang string, targetLang string) string {

	// translate
	var result string
	if constUseExamples {
		if sourceLang != "en" {
			panic(sourceLang)
		}
		result = translations[word][targetLang]
	} else {
		result = translateQuery(word, sourceLang, targetLang)
	}

	result = strings.ToLower(result)

	// remove accents
	// https://twinnation.org/articles/33/remove-accents-from-characters-in-go
	var t = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	var resultNoAccents, _, err = transform.String(t, result)
	if err != nil {
		panic(err)
	}

	// replace characters
	var resultFinal string
	var mTranslateChars2 = mTranslateCharsByLanguage[targetLang]
	for _, r := range resultNoAccents {
		// TODO support several characters in mTranslateChars/mTranslateCharsByLanguage as keys
		var c = string(r)
		if cTr, ok := mTranslateChars2[c]; ok {
			c = cTr
		}
		if cTr, ok := mTranslateChars[c]; ok {
			c = cTr
		}
		resultFinal += c
	}

	return resultFinal
}

//------------------------------------------------------------------------------
