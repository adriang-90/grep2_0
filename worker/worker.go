package worker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// this package handles the results

type Result struct {
	line       string
	LineNumber int
	Path       string
}

type Results struct {
	Inner []Result
}

func NewResult(line string, lineNumber int, Path string) Result {
	return Result{line, lineNumber, Path}
}

func FindInFile(path string, find string) *Results {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	results := Results{make([]Result, 0)}

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), find) {
			r := NewResult(scanner.Text(), lineNumber, path)
			results.Inner = append(results.Inner, r)
		}
		lineNumber += 1
	}
	if len(results.Inner) == 0 {
		return nil
	} else {
		return &results
	}
}
