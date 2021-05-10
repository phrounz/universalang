package translator_anylang

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

//------------------------------------------------------------------------------

type dictionary struct {
	translations map[string]string
}

func newDictionary() dictionary {
	return dictionary{translations: make(map[string]string)}
}

//------------------------------------------------------------------------------

var offlineDictionaries = make(map[string]dictionary)

const debugWordInit = "" //"understand"

//------------------------------------------------------------------------------

func Init(debugVerbose bool) {
	var files, err = filepath.Glob("data/*.txt")
	if err != nil {
		panic(err)
	}

	for _, filename := range files {

		var reFilepath = regexp.MustCompile("en-([^-]+)-enwiktionary.txt")
		var matchFilepath = reFilepath.FindStringSubmatch(filepath.Base(filename))
		if len(matchFilepath) > 0 {
			var language = matchFilepath[1]
			var dict = readDictionaryFile(filename)
			if language == "cmn" {
				language = "zh"
			}
			if language == "arb" {
				language = "ar"
			}
			offlineDictionaries[language] = dict
			if debugVerbose {
				fmt.Printf("  Loaded %s\n", language)
			}
		}
	}
}

//------------------------------------------------------------------------------

func readDictionaryFile(filename string) dictionary {
	var basename = filepath.Base(filename)
	var fd, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = fd.Close()
		if err != nil {
			panic(err)
		}
	}()
	var scanner = bufio.NewScanner(fd)

	var dict = newDictionary()

	var re = regexp.MustCompile("^(.+) ::(.*)$")
	// TODO "SEE: XXXX"

	for scanner.Scan() {
		var line = scanner.Text()
		if line[0:2] != "# " {
			var match = re.FindStringSubmatch(line)
			if len(match) > 0 {
				var word, finalTranslations = readDictionaryFileLine(match[1], match[2])
				if _, ok := dict.translations[word]; ok {
					dict.translations[word] += "," + strings.Join(finalTranslations, ",")
				} else {
					dict.translations[word] = strings.Join(finalTranslations, ",")
				}
			} else if line != "]] ::" {
				panic(line + " " + basename)
			}
		}
	}
	return dict
}

//------------------------------------------------------------------------------

func readDictionaryFileLine(word string, translationStr string) (string, []string) {

	if debugWordInit != "" && len(word) >= len(debugWordInit) && word[:len(debugWordInit)] == debugWordInit {
		fmt.Printf("===> %s\n", word)
	}
	word = removeInfo(word, "[", "]")
	word = removeInfo(word, "{", "}")
	word = removeInfo(word, "(", ")")
	word = removeInfo(word, "/", "/")
	word = strings.TrimSpace(word)
	word = strings.ToLower(word)

	if len(translationStr) > 0 && translationStr[0] == ' ' {
		translationStr = translationStr[1:]
	}
	var translations = doCommaSeparation(translationStr)

	if debugWordInit != "" && word == debugWordInit {
		fmt.Printf("%s :: %s\n", word, translationStr)
	}

	// TODO: "自然語言, 自然语言 /zìrányǔyán/" shall not take 自然語言
	var finalTranslations = make([]string, 0, len(translations))
	var finalTranslations2 = make([]string, 0, len(translations))
	var hasOccidentalAlphabet = false
	for _, translation := range translations {
		var tr string
		var partT = strings.Split(translation, "/")
		if len(partT) == 3 { // "自然语言 /zìrányǔyán/"
			tr = partT[1] // => "zìrányǔyán"
			hasOccidentalAlphabet = true
			if strings.Index(tr, ", ") != -1 { // 精錬所 /せいれんしょ, seirensho/
				tr = tr[strings.Index(tr, ", ")+2:]
			}
		} else if len(partT) == 1 {
			tr = translation
		}
		tr = removeInfo(tr, "[[[", "]]]")
		tr = removeInfo(tr, "[", "]")
		tr = removeInfo(tr, "{", "}")
		tr = removeInfo(tr, "{", "}")
		tr = removeSubstr(tr, " + comp.")
		tr = strings.TrimSpace(tr)
		tr = strings.ToLower(tr)
		var indexPolite = strings.Index(tr, "; [polite] ")
		if indexPolite != -1 {
			tr = tr[:indexPolite]
		}

		var reNotUsed = regexp.MustCompile("^not used in [a-z]+$")
		if !reNotUsed.MatchString(tr) {
			finalTranslations = append(finalTranslations, tr)
			if hasOccidentalAlphabet {
				finalTranslations2 = append(finalTranslations2, tr)
			}
		}
		if hasOccidentalAlphabet {
			finalTranslations = finalTranslations2
		}
	}

	if debugWordInit != "" && word == debugWordInit {
		fmt.Printf("%s :: %q\n", word, finalTranslations)
	}

	return word, finalTranslations
}

//------------------------------------------------------------------------------

func doCommaSeparation(translationStr string) (translations []string) {
	var commaSeparatorIndexes = findCommaSeparatorIndexes(translationStr)
	if len(commaSeparatorIndexes) == 0 {
		translations = []string{translationStr}
	} else {
		var iPrev = -1
		for _, i := range commaSeparatorIndexes {
			translations = append(translations, translationStr[iPrev+1:i])
			iPrev = i
		}
		translations = append(translations, translationStr[iPrev+1:])
	}
	return
}

//------------------------------------------------------------------------------

func findCommaSeparatorIndexes(str string) (out []int) {
	var withinC = ' '
	for i, c := range str {
		switch c {
		case '[':
			withinC = '['
		case '/':
			if withinC == ' ' {
				withinC = '/'
			} else if withinC == '/' {
				withinC = ' '
			} else {
				// TODO "آمن /ʾāmana/ (present tense: يؤمن /yúʾmin/)"
				//panic(str + " :: " + string(withinC))
			}
		case ']':
			if withinC == '[' {
				withinC = ' '
			}
		case '{':
			withinC = '{'
		case '}':
			if withinC == '{' {
				withinC = ' '
			}
		case '(':
			withinC = ')'
		case ')':
			if withinC == ')' {
				withinC = ' '
			}
		case ',':
			if withinC == ' ' {
				out = append(out, i)
			}
		}
	}
	return
}

//------------------------------------------------------------------------------

func removeSubstr(str string, substr string) string {
	var i = strings.Index(str, substr)
	if i != -1 {
		return str[:i] + str[i+len(substr):]
	}
	return str
}

//------------------------------------------------------------------------------

func removeInfo(str string, beginChar string, endChar string) string {
	var indexBegin = strings.Index(str, beginChar)
	if indexBegin != -1 {
		var indexEnd = strings.Index(str[indexBegin+len(beginChar):], endChar)
		if indexEnd != -1 {
			indexEnd += indexBegin + len(beginChar)
		}
		return str[:indexBegin] + str[indexEnd+len(endChar):]
	}
	return str
}

//------------------------------------------------------------------------------

func translateFromEnglishUsingOfflineDictionary(word string, targetLang string) string {
	if dict, ok1 := offlineDictionaries[targetLang]; ok1 {
		if translation, ok2 := dict.translations[word]; ok2 {
			return strings.Split(translation, ",")[0] // TODO choosing [0] is arbitrary
		}
		//panic(word + " - " + targetLang)
		//fmt.Printf("WARNING: %s - %s\n", word, targetLang)
		return ""
	}
	panic(targetLang)
}

//------------------------------------------------------------------------------
