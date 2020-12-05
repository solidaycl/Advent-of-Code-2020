package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strconv"
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
			} else {
				println("INVALID!")
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
	//fields := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	s = strings.ReplaceAll(s[1:], " ", "\", \"")
	s = "{\"" + strings.ReplaceAll(s, ":", "\":\"") + "\"}"
	var passport map[string]string

	// validate
	json.Unmarshal([]byte(s), &passport)

	println("byr: " + passport["byr"])
	byr, err := strconv.Atoi(passport["byr"])
	if !(err == nil && byr >= 1920 && byr <= 2002) {
		return false
	}

	println("iyr: " + passport["iyr"])
	iyr, err := strconv.Atoi(passport["iyr"])
	if !(err == nil && iyr >= 2010 && iyr <= 2020) {
		return false
	}

	println("eyr: " + passport["eyr"])
	eyr, err := strconv.Atoi(passport["eyr"])
	if !(err == nil && eyr >= 2020 && eyr <= 2030) {
		return false
	}

	println("hgt: " + passport["hgt"])
	if !isValidHeight(passport["hgt"]) {
		return false
	}

	println("hcl: " + passport["hcl"])
	if match, _ := regexp.MatchString("^#([0-9a-f]){6}$", passport["hcl"]); !match {
		return false
	}

	println("ecl: " + passport["ecl"])
	if match, _ := regexp.MatchString("^amb|blu|brn|gry|grn|hzl|oth$", passport["ecl"]); !match {
		return false
	}

	println("pid: " + passport["pid"])
	if match, _ := regexp.MatchString("^[0-9]{9}$", passport["pid"]); !match {
		return false
	}

	return true
}

func isValidHeight(s string) bool {
	if len(s) <= 3 {
		return false
	}
	hgt, _ := strconv.Atoi(s[:len(s)-2])
	if match, _ := regexp.MatchString("([0-9]+)cm", s); match {
		return hgt >= 150 && hgt <= 193
	}
	if match, _ := regexp.MatchString("([0-9]+)in", s); match {
		return hgt >= 59 && hgt <= 76
	}
	return false
}
