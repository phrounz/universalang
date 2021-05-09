package translator_anylang_simplified

//------------------------------------------------------------------------------

// https://en.wikipedia.org/wiki/International_Phonetic_Alphabet
//
// This map translates alphabets into IPA.
//
// TODO manage altenatives e.g. "ø" in danish should give "ø" or "œ", or "ɶ", depending of the word.
var mTranslateCharsByLanguage = map[string]map[string]string{
	// https://en.wikipedia.org/wiki/Help:IPA/Mandarin (pinyin)
	"zh": map[string]string{
		"x":  "ɕ",
		"f":  "f",
		"y":  "j",
		"i":  "j",
		"g":  "k",
		"k":  "kʰ",
		"m":  "m",
		"n":  "n",
		"ŋ":  "ŋ",
		"b":  "p",
		"p":  "r",
		"r":  "ʐ", // ɻ i
		"s":  "s",
		"sh": "ʂ",
		"d":  "t",
		"t":  "tʰ",
		"j":  "tɕ",
		"q":  "tɕʰ",
		"z":  "ts",
		"c":  "tsʰ",
		"zh": "ʈʂ",
		"ch": "ʈʂʰ",
		"w":  "w",
		"u":  "w",
		"h":  "x",
		"yu": "ɥ",
		"a":  "a",
		"e":  "e", //ɛ~æ/ə/ɤ
		"er": "ɚ",
		"yi": "i",
		"o":  "o",
		//"i":  "ɻ̩", //
		//"ʊ":
		"ai": "ai",
		"ao": "au",
		"ei": "ei",
		"ou": "ou",
		"ā":  "á",
		"á":  "ǎ",
		"à":  "â",
		// TODO other vovels with accent
		// "a": "a",
	},
	// https://en.wikipedia.org/wiki/Help:IPA/French
	"fr": map[string]string{
		"c": "k",
		"ç": "s",
		// TODO
	},
	// https://en.wikipedia.org/wiki/Help:IPA/Danish
	"da": map[string]string{
		"a":  "ɑ",  //ɑː/æ/ɛː
		"å":  "ʌ",  //ɔː
		"or": "ɒ",  //ɒː
		"æ":  "æː", //eː/ɛ
		"e":  "e",  //e̝/e̝ː/e/ə
		"i":  "i",  //i/iː
		"o":  "o",  //o/o:/ɔ/
		"ø":  "ø",  // øː/œ/œː/ɶ/ɶː
		"u":  "u",  //u/uː
		"y":  "y",  //yː
		"er": "ɐ",
		"et": "ð̩",
		"el": "l̩",
		"en": "m̩", //n̩/ŋ̍
		"sj": "ɕ",
		"d":  "ð", //t
		"f":  "f",
		"h":  "h",
		"j":  "j",
		"g":  "k", //j
		"k":  "kʰ",
		"l":  "l",
		"m":  "m",
		"n":  "n",
		"ng": "ŋ",
		"b":  "p",
		"p":  "pʰ",
		"r":  "ʁ", //ɐ̯
		"s":  "s",
		"t":  "tsʰ",
		"tj": "tɕ",
		"v":  "v", //w
		// TODO
	},
	// TODO other languages
}

//------------------------------------------------------------------------------
