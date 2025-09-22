package main

import (
	"6.5840/mr"
	"strings"
	"unicode"
	"fmt"
	"sort"
)

func Map(document string, value string) (res []mr.KeyValue) {
	m := make(map[string]bool)
	words := strings.FieldsFunc(value, func(r rune) bool { return !unicode.IsLetter(r) })
	for _, w := words {
		m[w] = true
	}
	for w:= range m{
		kv := mr.KeyValue{w, document}
		res = append(res, kv)
	}
	return res
}

func Reduce(key string, values []string) string {
	sort.Strings(values)
	return fmt.Sprintf("%d %s", len(values), strings.Join(values, ","))
}