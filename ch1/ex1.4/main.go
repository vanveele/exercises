// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	file_list := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, file_list)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, file_list)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d in %s\t%s\n", n, strings.Join(file_list[line], ","), line[:8]+"... ")
		}
	}
}

func countLines(f *os.File, counts map[string]int, files map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		files[input.Text()] = appendIfMissing(files[input.Text()], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}

func appendIfMissing(slice []string, i string) []string {
	for _, name := range slice {
		if name == i {
			return slice
		}
	}
	return append(slice, i)
}
