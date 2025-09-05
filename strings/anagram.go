package main

import (
	"slices"
	"unicode"
)

func Anagram(s1, s2 string) bool {
	s1 = _normalize(s1)
	s2 = _normalize(s2)
	return true
}

func _normalize(s string) string {
	var r []rune
	for _, b := range s {
		b = unicode.ToLower(b)
		if b >= 'a' && b <= 'z' {
			r = append(r, b)
		}
	}
	slices.Sort(r)
	return string(r)
}
