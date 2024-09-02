package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "Usage: mygrep -E <pattern>\n")
		os.Exit(2)
	}

	pattern := os.Args[2]

	fmt.Printf("Pattern is %q\n", pattern)

	line, err := io.ReadAll(os.Stdin) // Assuming there is only one line
	if err != nil {
		fmt.Fprintf(os.Stderr, "Usage: mygrep -E <pattern>\n")
		os.Exit(2)
	}

	fmt.Printf("Line is %q\n", line)

	ok, err := matchLine(line, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !ok {
		os.Exit(1)
	} else {
		fmt.Println("Hooyaa..!")
	}
}

func matchLine(line []byte, pattern string) (bool, error) {
	patternLength := utf8.RuneCountInString(pattern)
	fmt.Println("RuneCountInString vs Len", patternLength, len(pattern))
	if patternLength == 0 {
		return false, fmt.Errorf("unsupported pattern: %q", pattern)
	}

	var ok bool

	if pattern == "\\d" {
		ok = bytes.ContainsAny(line, "1234567890")
	} else if pattern == "\\w" {
		ok = bytes.ContainsAny(line, "abcdefghijklmnopqrstvuwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")
	} else if patternLength == 1 {
		ok = bytes.ContainsAny(line, pattern)
	} else if pattern[0] == '[' && pattern[len(pattern)-1] == ']' {
		if pattern[1] == '^' {
			ok = !bytes.ContainsAny(line, pattern[2:len(pattern)-1])
		} else {
			ok = bytes.ContainsAny(line, pattern[1:len(pattern)-1])
		}
	}

	return ok, nil
}
