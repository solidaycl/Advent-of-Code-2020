package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2.input")

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

	sum := 0
	for _, x := range txtlines {
		parts := strings.Split(x, ": ")
		password := parts[1]
		parts = strings.Split(parts[0], " ")
		c := parts[1]
		parts = strings.Split(parts[0], "-")
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])

		if strings.Count(password, c) >= min && strings.Count(password, c) <= max {
			sum++;
		}
	}
	println(sum)
}
