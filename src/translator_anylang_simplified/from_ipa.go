package translator_anylang_simplified

//------------------------------------------------------------------------------

// https://en.wikipedia.org/wiki/International_Phonetic_Alphabet
//
// This map translates IPA into a readable english alphabet.
//
// https://www.speechactive.com/english-vowels-ipa-international-phonetic-alphabet/
//
// TODO return international phonetic alphabet as output instead of trying to eliminate IPA characters?
var mTranslateCharsFinal = map[string]string{
	"ʰ":  "",
	" ":  "",
	"ˈ":  "",
	"'":  "", // (that's not the same than above, haha)
	"’":  "",
	"ʁ":  "r",
	"ɕ":  "sh",
	"ŋ":  "ng",
	"ə":  "e",
	"ä":  "o",
	"ē":  "i",
	"ø":  "o",
	"œ":  "e",
	"ɥ":  "u",
	"ʈʂ": "tch",
	"ʈs": "tch",
	"ʈ":  "t",
	"ʂ":  "sh",
	"ɑ":  "a",
	"ð":  "th",
	"m̩": "m",
	"n̩": "n",
	"h̩": "h", //?
	"ō":  "o",
	"ː":  "",
	"æ":  "a",
	"ɐ":  "e",
	"ʾ":  "",
	"ɣ":  "g", //between 'h'ome and 'g'oo
	"ʲ":  "",
	"ɨ":  "i",
	"ʊ":  "e",
	"ɛ":  "e",
	"ʌ":  "e",
	"ɒ":  "o", //e?
	"ʕ":  "a", // re https://www.youtube.com/watch?v=3057MbWmH1k
	":":  "",
	"ʃ":  "ch",
}

//------------------------------------------------------------------------------
