package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
