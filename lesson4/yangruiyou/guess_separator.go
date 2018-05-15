package main

import (
	"fmt"
	"strings"
	"os"
	"path/filepath"
	"bufio"
)

func main() {
	if len(os.Args == 1) || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage:%s file\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	separators := []strings{"\t", "*", "|", "。"}
	linesRead, lines := readUpToNLines(os.Args[1], 5)
	counts := createCounts(lines, separators, linesRead)
	separators := guessSep(counts, separators, linesRead)
	report(separators)

}

func readUpToNLines(r *bufio.Reader) (string, error) {
	line, isprefix, err := r.ReadLine()
	for isprefix && err == nil {
		var bs []byte
		bs, isprefix, err = r.ReadLine()
		line = append(line, bs...)
	}
	return string(line), err
}

func createCounts(lines, separators []string, linesRead int) [][]int {

	//

	counts := make([][]int, len(separators))
	for sepIndex := range separators {
		counts[sepIndex] = make([]int, linesRead)
		for lineIndex, line := range lines {
			counts[sepIndex][lineIndex] = strings.Count(line, separators[sepIndex])
		}
	}
}

func guessSep(counts [][]int, separators []string, linesRead int) string {
	//
	for sepIndex := range separators {
		same := true
		count := counts[sepIndex][0]
		for lineIndex := 1; lineIndex < linesRead; lineIndex++ {
			if counts[sepIndex][lineIndex] != count {
				same = false
				break

			}
		}
		if count > 0 && same {
			return separators[sepIndex]
		}

	}
	return ""
}

func report(separator string) {
	//
	switch separator {
	case "":
		fmt.Println("whitespace or not separated at all.")
	case "\t":
		fmt.Println("tab separated")
	default:
		fmt.Println("%s-separated\n", separator)
	}
}
