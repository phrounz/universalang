package main

//------------------------------------------------------------------------------

const constUseExamples = true

//------------------------------------------------------------------------------

var exampleTranslations = map[string]map[string]string{
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
		"ja": "cha", //ocha
		"ar": "shay",
		"ru": "chay",
	},
	"yes": map[string]string{
		"en": "yes", //?
		"fr": "oui",
		"es": "sí",
		"zh": "shì",
		"da": "ja",
		"id": "iya",
		"ja": "hai",
		"ar": "nem",
		"ru": "da",
	},
	"no": map[string]string{
		"en": "nō",
		"fr": "non",
		"es": "no",
		"zh": "bù",
		"da": "ingen",
		"id": "tidak",
		"ja": "iie", //bango?
		"ar": "la",
		"ru": "net",
	},
	"love": map[string]string{
		"en": "ləv",
		"fr": "aimer",
		"es": "amar",
		"zh": "qù ài",
		"da": "elske",
		"id": "mencintai",
		"ja": "aisuru",
		"ar": "yuhibu",
		"ru": "lyubit'",
	},
}

//------------------------------------------------------------------------------
