package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	fmt.Println("hello world !")
	combineOxford3000and5000()
	fmt.Println("finsihed")
}

func checkDuplicateBetweenOxford3000And5000() {
	fmt.Println("starting app...")

	// read files
	lines3000 := collectLines("oxford_3000.txt")
	lines5000 := collectLines("oxford_5000.txt")

	// lets cache 3000 first
	m := make(map[string]bool)
	for _, e := range *lines3000 {
		m[e] = true
	}

	// lets check duplicates in 5000
	count := 0
	for _, e := range *lines5000 {
		if m[e] {
			fmt.Printf("found duplicate %s\n", e)
			count++
		}
	}

	// stdout
	fmt.Printf("duplicate count: %d\n", count)
}

func combineOxford3000and5000() {
	fmt.Println("starting app...")

	// read files
	lines3000 := collectLines("oxford_3000.txt")
	lines5000 := collectLines("oxford_5000.txt")

	// filter unique words
	m := make(map[string]bool)
	for _, e := range *lines3000 {
		m[e] = true
	}
	for _, e := range *lines5000 {
		m[e] = true
	}

	// put in slice to be sorted later
	s := make([]string, 0, len(m))
	for word := range m {
		s = append(s, word)
	}

	// sort
	sort.Strings(s)

	// Open the file in append mode, create if not exists
	file, err := os.OpenFile("oxford_5000_sorted.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// append to output file
	for _, word := range s {
		if _, err := file.WriteString(word + "\n"); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func collectLines(file string) *[]string {
	// read file
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error when opening file: %s", err)
		os.Exit(1)
	}
	reader := bufio.NewReader(f)

	lines := []string{}

	// read line by line
	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}

		// iterate until all chars in a line is consumed
		for isPrefix {
			nextPart, nextIsPrefix, err := reader.ReadLine()
			if err != nil {
				break
			}
			line = append(line, nextPart...)
			isPrefix = nextIsPrefix
		}
		lines = append(lines, string(line))
	}

	return &lines
}
