package main

import (
	"errors"
	"fmt"
)

type wordType []rune
type listType []wordType

var empty = wordType("")

func main() {
	fmt.Println(listType{})
	fmt.Println(listOfStringToListType([]string{}))
}

func POS(word wordType, n int) wordType {
	if n >= LENGTH(word) {
		return wordType([]rune{})
	}
	counter := 0

	for counter != n {
		word = word[1:]
		counter++
	}
	return wordType{word[0]}
}

func listOfStringToListType(list []string) listType {
	var newList listType
	for _, word := range list {
		newList = append(newList, wordType(word))
	}
	if len(list) == 0 {
		newList = []wordType{}
	}
	return newList
}

func LENGTH(word wordType) int {
	length := 0
	if Equal(word, empty) {
		return length
	}
	length++

	for !Equal(TAIL(word), empty) {
		word = word[1:]
		length++
	}
	return length
}

func TAIL[T wordType | rune](list []T) []T {
	return list[1:]
}

func Equal(x, y wordType) bool {
	if len(x) != len(y) {
		return false
	}
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func MAKELIST[T wordType | rune](element T, list []T) []T {
	OUT := append([]T{element}, list...)
	return OUT
}

func REV(word wordType) wordType {
	if Equal(word, empty) {
		return empty
	}
	OUT := wordType{}
	for i := 0; i < LENGTH(word); i++ {
		OUT = MAKELIST(word[i], OUT)

	}
	return OUT
}

func CON(P wordType, Q wordType) wordType {
	if Equal(Q, empty) {
		return P
	}
	if Equal(P, empty) {
		return Q
	}
	P = REV(P)
	for i := 0; i < LENGTH(Q); i++ {
		P = MAKELIST(Q[i], P)
	}

	return REV(P)
}

func SUBSTRING(word wordType, looking wordType) bool {
	if Equal(looking, empty) {
		return true
	}
	if Equal(word, empty) {
		return false
	}
	i := 0
	for i <= LENGTH(word)-LENGTH(looking) {
		found := true
		for j := 0; j <= LENGTH(looking)-1; j++ {
			if word[i+j] != looking[j] {
				found = false
				break
			}
		}
		if found {
			return true
		}
		i++
	}
	return false
}

func REMOVESYMBOL(word wordType, i int) (wordType, error) {
	if Equal(word, empty) {
		return empty, errors.New("nie da się usunąć symbolu - P jest puste")
	}
	if i >= LENGTH(word) {
		return empty, errors.New("nie da się usunąć symbolu - i jest za duże")
	}

	frontP := wordType{}
	for j := 0; j < i; j++ {
		frontP = CON(frontP, wordType{word[0]})
		word = TAIL(word)
	}
	word = CON(frontP, TAIL(word))
	return word, nil
}

func EQWORD(word1 wordType, word2 wordType) bool {
	if LENGTH(word1) != LENGTH(word2) {
		return false
	}
	return SUBSTRING(word1, word2)
}

func langIsEmpty(language listType) bool {
	return len(language) == 0
}

func MEMBERSHIP(language listType, word wordType) bool {
	for !langIsEmpty(language) {
		if EQWORD(language[0], word) {
			return true
		}
		language = TAIL(language)
	}
	return false
}

func REMOVE(language listType, word wordType) listType {
	endList := listType{}
	for !langIsEmpty(language) {
		if !EQWORD(language[0], word) {
			endList = MAKELIST(language[0], endList)
		}
		language = TAIL(language)
	}
	return endList
}

func REMOVEREPETITIONS(language listType) listType {
	endList := listType{}
	for !langIsEmpty(language) {
		if !MEMBERSHIP(endList, language[0]) {
			endList = MAKELIST(language[0], endList)
		}
		language = TAIL(language)
	}
	return endList
}

func EQLANG(language1 listType, language2 listType) bool {
	for !langIsEmpty(language1) {
		if !MEMBERSHIP(language2, language1[0]) {
			return false
		}
		language2 = REMOVE(language2, language1[0])
		language1 = TAIL(language1)
	}
	return langIsEmpty(language2)
}

func CONLANG(language1 listType, language2 listType) listType {
	endList := listType{}
	for !langIsEmpty(language1) {
		lang2save := language2
		for !langIsEmpty(language2) {
			endList = MAKELIST(CON(language1[0], language2[0]), endList)
			language2 = REMOVE(language2, language2[0])
		}
		language2 = lang2save
		language1 = REMOVE(language1, language1[0])
	}
	return REMOVEREPETITIONS(endList)
}

func POWLANG(language listType, n int) listType {
	if n == 0 {
		return listType{}
	}
	if n == 1 {
		return language
	}
	return CONLANG(POWLANG(language, n-1), language)
}

func UNION(language1 listType, language2 listType) listType {
	for !langIsEmpty(language2) {
		language1 = MAKELIST(language2[0], language1)
		language2 = TAIL(language2)
	}
	return REMOVEREPETITIONS(language1)
}

func MEET(language1 listType, language2 listType) listType {
	endList := listType{}
	for !langIsEmpty(language1) {
		if MEMBERSHIP(language2, language1[0]) {
			endList = MAKELIST(language1[0], endList)
		}
		language1 = TAIL(language1)
	}
	return endList
}

func SETMINUS(language1 listType, language2 listType) listType {
	endList := listType{}
	for !langIsEmpty(language1) {
		if !MEMBERSHIP(language2, language1[0]) {
			endList = MAKELIST(language1[0], endList)
		}
		language1 = TAIL(language1)
	}
	return endList
}

func REVLANG(language listType) listType {
	endList := listType{}
	for !langIsEmpty(language) {
		endList = MAKELIST(REV(language[0]), endList)
		language = TAIL(language)
	}
	return endList
}

func SUBWORDS(word wordType, length int) listType {
	endList := listType{}
	for i := 0; i <= LENGTH(word)-length; i++ {
		if length < 1 {
			break
		}
		subword := wordType{}
		for j := 0; j <= length-1; j++ {
			subword = CON(subword, POS(word, i+j))
		}
		endList = MAKELIST(wordType(subword), endList)
	}

	return REMOVEREPETITIONS(endList)
}

func ALLSUBWORDS(word wordType) listType {
	endList := listType{}
	for i := 1; i <= LENGTH(word); i++ {
		endList = UNION(endList, SUBWORDS(word, i))
	}
	return endList
}
