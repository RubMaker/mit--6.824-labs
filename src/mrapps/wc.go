package main

import (
	"strconv"
	"strings"
	"unicode"

	"6.5840/mr"
)

func Map(filename string, contents string) []mr.KeyValue {
	ff := func(r rune) bool {
		return !unicode.IsLetter(r)
	}
	// Split contents into words.
	words := strings.FieldsFunc(contents, ff)
	var kva []mr.KeyValue
	for _, word := range words {
		kv := mr.KeyValue{word, "1"}
		kva = append(kva, kv)
	}
	return kva
}

func Reduce(key string, values []string) string {
	// Return count of occurrences as a string.
	return strconv.Itoa(len(values))
}
