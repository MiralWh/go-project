package helperstrings

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func GetStringSize(value string) int {
	return utf8.RuneCountInString(value)
}

type RuneSize struct {
	Char rune
	Size int
}

func GetRuneCountInString(value string, runes []rune) (bool, []RuneSize) {

	size := GetStringSize(value)

	if size == 0 || len(runes) == 0 {
		return false, make([]RuneSize, 0)
	}

	array := make([]RuneSize, len(runes))

	for i := range len(runes) {

		item := runes[i]
		substr := string(item)

		array[i] = RuneSize{Char: item, Size: strings.Count(value, substr)}
	}

	return true, array
}

func CapitalizeWords(value string) (bool, string) {

	size := GetStringSize(value)

	if size == 0 {
		return false, value
	}

	var b strings.Builder

	words := strings.Fields(value)

	for i, s := range words {

		if s == "" {
			continue
		}

		r, _ := utf8.DecodeRuneInString(s)

		runes := []rune(s)

		b.Reset()

		b.WriteString(string(unicode.ToUpper(r)))
		b.WriteString(string(runes[1:]))

		words[i] = b.String()
	}

	return true, strings.Join(words, " ")
}

func СheckFormulaBrackets(s string) (bool, int, int) {

	openCount, closeCount, all := 0, 0, 0

	for _, char := range s {

		switch char {
		case '(':
			openCount++
			all++
		case ')':
			closeCount++
			all--
		}
	}

	return all == 0, openCount, closeCount
}
