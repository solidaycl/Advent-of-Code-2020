package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day4.input")

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

	numValidPassports := 0
	s := ""
	for _, line := range txtlines {
		if line == "" {
			if isValid(s) {
				numValidPassports++
			}
			s = ""
		} else {
			s += " " + line
		}
	}
	println(numValidPassports)
}

func isValid(s string) bool {

	//jsonify
	fields := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	s = strings.ReplaceAll(s[1:], " ", "\", \"")
	s = "{\"" + strings.ReplaceAll(s, ":", "\":\"") + "\"}"
	passport := make(map[string]interface{})

	// validate
	json.Unmarshal([]byte(s), &passport)
	for _, field := range fields {
		if passport[field] == nil {
			return false
		}
	}
	return true
}
