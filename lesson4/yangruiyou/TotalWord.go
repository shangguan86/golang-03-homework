package main

import (
	"runtime"
	"path/filepath"
	"os"
	"log"
	"bufio"
	"unicode"
	"strings"
	"unicode/utf8"
	"fmt"
	"sort"
)

func main() {
	//统计单次的频次
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage:%s <file> [<file2> [... <fileN>]]\n ", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	frequecyForWord := map[string]int{}
	for _, filename := range commandLineFiles(os.Args[1:]) {
		reportBywords(frequecyForWord)
		wordsForFrequency := invertStringInMap(frequecyForWord)
		reportByFrequency(wordsForFrequency)
	}

}

func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name)

			} else if matches != nil {
				args = append(args, matches...)
			}
		}
		return args

	}
	return files

}

func updateFrequencies(filename string, freqencyForWord map[string]int) {
	var file *os.File
	var err error
	if file, err = os.Open(filename); err != nil {
		log.Println("failed to open the file: ", err)

	}
	defer file.Close()
	readAndUpdateFrequencies(bufio.NewReader(file), freqencyForWord)

}

func readAndUpdateFrequencies(reader *bufio.Reader, frequencyForWord map[string]int) {

}

func SplitOnNonLetters(s string) []string {
	noALetter := func(char rune) bool {
		return !unicode.IsLetter(char)

	}
	return strings.FieldsFunc(s, noALetter)
}

func reportBywords(frequencyForWord map[string]int) {
	words := make([]string, 0, len(frequencyForWord))
	wordWidth, frequencyWidth := 0, 0
	for word, frequency := range frequencyForWord {
		words := append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}
	sort.Strings(words)
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth, frequencyForWord[word])
	}
}

func invertStringInMap(intForString map[string]int) map[int][]string {
	stringsForInt := make(map[int][]string, len(intForString))
	for key, value := range intForString {
		stringsForInt[value] = append(stringsForInt[value], key)
	}
	return stringsForInt

}

func reportByFrequency(wordsForFrequency map[int]string) {
	frequencies := make([]int, 0, len(wordsForFrequency))
	for frequency := range wordsForFrequency {
		frequencies = append(frequencies, frequency)
	}
	sort.Ints(frequencies)
	width := len(fmt.Sprint(frequencies[len(frequencies)-1]))
	fmt.Println("Frequency->Words")
	for _, frequency := range frequencies {
		words := wordsForFrequency[frequency]
		sort.Strings(words)
		fmt.Printf("%*d %s\n", width, frequency, strings.Join(words, ", "))
	}

}
