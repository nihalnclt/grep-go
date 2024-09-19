package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "Usage: mygrep -E <pattern>\n")
		os.Exit(2)
	}

	pattern := os.Args[2]

	line, err := io.ReadAll(os.Stdin) // Assuming there is only one line
	if err != nil {
		fmt.Fprintf(os.Stderr, "Usage: mygrep -E <pattern>\n")
		os.Exit(2)
	}

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
	patternLength := len(pattern)
	if patternLength == 0 {
		return false, fmt.Errorf("unsupported pattern: %q", pattern)
	}

	if pattern[0] == '^' {
		return matchPattern(line, pattern[1:], 0), nil
	}

	for i := 0; i < len(line); i++ {
		fmt.Println("RUNNING MATCH LINE LOOP", string(line[i]))
		if matchPattern(line, pattern, i) {
			return true, nil
		}
	}

	return false, nil
}

func matchPattern(line []byte, pattern string, start int) bool {
	for i := 0; i < len(pattern); i++ {
		fmt.Println("---PATTERN LOOP")
		fmt.Printf("START IS %d - LEN(LINE) is %d\n", start, len(line))
		if start >= len(line)-1 {
			return pattern[i] == '$'
		}
		if pattern[i] == '\\' && i+1 < len(pattern) {
			if pattern[i+1] == 'd' && !unicode.IsDigit(rune(line[start])) {
				return false
			} else if pattern[i+1] == 'w' && !unicode.IsLetter(rune(start)) {
				return false
			} else {
				i++
			}
		} else if pattern[i] == '[' && pattern[i+1] == '^' {
			endPos := strings.Index(pattern[i:], "]")
			if strings.Contains(pattern, pattern[i+1:endPos]) {
				return false
			}
			i = endPos
		} else if pattern[i] == '[' {
			endPos := strings.Index(pattern[i:], "]")
			if !strings.Contains(pattern, pattern[i+1:endPos]) {
				return false
			}
			i = endPos
		} else {
			if start < len(line) && line[start] != pattern[i] {
				return false
			}
		}

		start++
	}

	return true
}
