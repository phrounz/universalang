package main

import (
	"fmt"
	"strings"
)

//------------------------------------------------------------------------------

const constDebug = true

//------------------------------------------------------------------------------

func translateEnglishWordToUniversalang(word string) string {

	var languages = []string{"en", "fr", "es", "zh", "da", "id", "ja", "ar"}

	var trByLanguage = make(map[string]string)
	for _, language := range languages {
		var result = translateWordSimplified(word, "en", language)
		// if isVowel([]rune(result)[0]) {
		// 	result = " " + result // add a space to keep consonants/vowels alignment TODO discuss
		// }
		trByLanguage[language] = result
	}
	//trByLanguage["en"] = word

	if constDebug {
		fmt.Printf("\n")
		for language, wordTr := range trByLanguage {
			fmt.Printf("%s %s\n", language, wordTr)
		}
	}

	var a = &aggregator{}
	a.prepare(trByLanguage)
	a.aggregate(trByLanguage)
	if constDebug {
		a.displayMap()
	}

	return a.finalize()
}

//------------------------------------------------------------------------------

type aggregator struct {
	nbBySoundByIndex []map[string]int
	maxLength        int
}

//------------------------------------------------------------------------------

func (a *aggregator) prepare(trByLanguage map[string]string) {

	a.maxLength = -1
	for _, wordTr := range trByLanguage {
		if len(wordTr) > a.maxLength {
			a.maxLength = len(wordTr)
		}
	}

	a.nbBySoundByIndex = make([]map[string]int, a.maxLength)
	for i := range a.nbBySoundByIndex {
		a.nbBySoundByIndex[i] = make(map[string]int)
	}
}

//------------------------------------------------------------------------------

func (a *aggregator) setSound(i int, sound string) {
	if _, ok := a.nbBySoundByIndex[i][sound]; !ok {
		a.nbBySoundByIndex[i][sound] = 0
	}
	a.nbBySoundByIndex[i][sound]++
}

//------------------------------------------------------------------------------

func (a *aggregator) aggregate(trByLanguage map[string]string) {
	for _, wordTr := range trByLanguage {
		var i = 0
		var sound string
		var c rune
		var isVowelCurrent = false
		var firstRune = true
		for _, c = range wordTr {
			var isVo = isVowel(c)
			// TODO add semivowels
			if firstRune {
				isVowelCurrent = isVo
			}
			if isVo && isVowelCurrent {
				sound += string(c)
			} else if !isVo && !isVowelCurrent {
				sound += string(c)
			} else {
				isVowelCurrent = !isVowelCurrent
				if len(sound) == 0 {
					panic(sound)
				}
				a.setSound(i, sound)
				sound = string(c)
				i++
			}
			firstRune = false
		}
		if len(sound) > 0 {
			a.setSound(i, sound)
			sound = ""
		}

		i++
		for ; i < a.maxLength; i++ {
			a.setSound(i, " ")
		}
	}

}

//------------------------------------------------------------------------------

func (a *aggregator) displayMap() {
	for i, nbByRunes := range a.nbBySoundByIndex {
		fmt.Printf("%d|| ", i)
		for rune, nb := range nbByRunes {
			fmt.Printf("%s:%d ", string(rune), nb)
		}
		fmt.Printf("\n")
	}
}

//------------------------------------------------------------------------------

func (a *aggregator) finalize() string {

	//---
	// algorithm to select best sounds

	var outputBest []string
	var outputNoBl []string

	for _, nbBySound := range a.nbBySoundByIndex {

		var maxNbBest = -1
		var maxNbNoBl = -1
		var maxSoundBest = "?"
		var maxSoundNoBl = "?"

		for sound, nb := range nbBySound {

			// TODO find similarities between nearby sounds

			if nb >= maxNbNoBl && sound != " " {
				maxNbNoBl = nb
				maxSoundNoBl = sound
				// TODO what to do if equality
			}

			if nb > maxNbBest {
				maxNbBest = nb
				maxSoundBest = sound
			} else if nb == maxNbBest /*&& sound != " " */ {
				maxSoundBest = sound
				// TODO what to do if equality
			}

		}
		outputBest = append(outputBest, maxSoundBest)
		outputNoBl = append(outputNoBl, maxSoundNoBl)
	}

	//---
	// algorithm to avoid several words like "foo b"
	// if outputBest is "foo a"
	// and outputNoBl is "foobaz"
	// then final output is "fooba"
	//
	// also reduces chances of not having any vowel in final output

	var firstBl = -1
	var lastNotBl = -1
	var noVowelYet = true
	for i := 0; i < len(outputBest); i++ {
		if outputBest[i] == " " {
			if firstBl == -1 {
				firstBl = i
			}
			if noVowelYet && outputNoBl[i] != " " && isVowel([]rune(outputNoBl[i])[0]) {
				lastNotBl = i + 1
				noVowelYet = false
			}
		} else {
			if isVowel([]rune(outputNoBl[i])[0]) {
				noVowelYet = false
			}
			lastNotBl = i + 1
		}
	}
	if firstBl == -1 {
		firstBl = len(outputBest)
	}
	if lastNotBl == -1 {
		panic("only spaces?")
	}

	if constDebug {
		fmt.Printf("===> %d %d %s %s\n", firstBl, lastNotBl, strings.Join(outputBest, "."), strings.Join(outputNoBl, "."))
	}
	return strings.Join(outputBest[0:firstBl], "") + strings.Join(outputNoBl[firstBl:lastNotBl], "")
}

//------------------------------------------------------------------------------

func isVowel(c rune) bool {
	switch c {
	case 'a':
		fallthrough
	case 'e':
		fallthrough
	case 'i':
		fallthrough
	case 'o':
		fallthrough
	case 'u':
		fallthrough
	case 'y':
		return true
	}
	return false
}

//------------------------------------------------------------------------------
