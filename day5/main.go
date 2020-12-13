package main

import (
	"bufio"
	"log"
	"os"
)

func getInput() []string {
	file, err := os.Open("input.txt")

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

	defer file.Close()
	return txtlines
}

func main() {
	println(part1())
	println(part2())
}
