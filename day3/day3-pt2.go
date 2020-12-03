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

	product := getTrees(1, 1, txtlines) * getTrees(3, 1, txtlines) * getTrees(5, 1, txtlines) * getTrees(7, 1, txtlines) * getTrees(1, 2, txtlines)
	println(product)
}

func getTrees(right int, down int, txtlines []string) int {

	numTrees := 0
	x := 0
	for i := 0; i < len(txtlines); i += down {
		if txtlines[i][x%len(txtlines[i])] == '#' {
			numTrees++
		}
		x = x + right
	}
	return numTrees
}
