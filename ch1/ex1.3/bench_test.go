// Echo2 prints its command-line arguments.
package main

import (
	"strings"
	"testing"
)

var args = []string{"sss", "uuu", "n", "set", "over", "the", "ocean"}

func appendStrings(input []string) {
	var s, sep string
	for i := 1; i < len(input); i++ {
		s += sep + input[i]
		sep = " "
	}
}

func appendStringsRange(input []string) {
	s, sep := "", ""
	for _, arg := range input {
		s += sep + arg
		sep = " "
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appendStrings(args)
	}
}

func BenchmarkAppendRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appendStringsRange(args)
	}
}
