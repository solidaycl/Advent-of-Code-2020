package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1.input")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []int

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		txtlines = append(txtlines, i)
	}

	file.Close()

	for _, i := range txtlines {
		for _, j := range txtlines {
			for _, k := range txtlines {
				if i+j+k == 2020 {
					fmt.Println(i * j * k)
				}
			}
		}
	}
}
