package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day3.input")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		i := scanner.Text()
		txtlines = append(txtlines, i)
	}

	file.Close()

	numTrees := 0
	x := 0
	for _, row := range txtlines {
		if row[x%len(row)] == '#' {
			numTrees++
		}
		x = x + 3
	}

	println(numTrees)
}
