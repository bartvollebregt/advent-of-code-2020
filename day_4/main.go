package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
)

var PART1_REQUIRED_FIELDS = [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

var PART2_REGEX_RULES = map[string]*regexp.Regexp{
	"byr": regexp.MustCompile(`^(19[2-8][0-9]|199[0-9]|200[0-2])$`),
	"iyr": regexp.MustCompile(`^(201[0-9]|2020)$`),
	"eyr": regexp.MustCompile(`^(202[0-9]|2030)$`),
	"hgt": regexp.MustCompile(`^((1[5-8][0-9]|19[0-3])cm)|((59|6[0-9]|7[0-6])in)$`),
	"hcl": regexp.MustCompile(`^#[0-9a-f]{6}$`),
	"ecl": regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`),
	"pid": regexp.MustCompile(`^\d{9}$`),
}

func readFileWithReadFile(filename string) (file string) {

	b, err := ioutil.ReadFile(filename) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

func mapLinesToPassports(file string) (passports []map[string]string) {
	var unparsedPassports = strings.Split(file, "\n\n")
	for _, unparsedPassport := range unparsedPassports {
		unparsedPassport = strings.Replace(unparsedPassport, "\n", " ", -1)
		var fields = strings.Split(unparsedPassport, " ")
		var passport = make(map[string]string)
		for _, field := range fields {
			var splittedValue = strings.Split(field, ":")
			var fieldName = splittedValue[0]
			var value = strings.TrimSpace(splittedValue[1])
			passport[fieldName] = value
		}
		passports = append(passports, passport)
	}

	return passports
}

func validatePassports(passports []map[string]string, validateRules bool) (validPassports []map[string]string) {
	for _, passport := range passports {
		var passportValid = true
		for _, fieldName := range PART1_REQUIRED_FIELDS {
			if _, ok := passport[fieldName]; !ok {
				passportValid = false
			} else if validateRules {
				var pattern = PART2_REGEX_RULES[fieldName]
				if !pattern.MatchString(passport[fieldName]) {
					passportValid = false
				}
			}
		}
		if passportValid {
			validPassports = append(validPassports, passport)
		}
	}
	return validPassports
}

func main() {
	start := time.Now()

	var lines = readFileWithReadFile("./day_4/input.txt")
	var passports = mapLinesToPassports(lines)

	var validPassportsPart1 = validatePassports(passports, false)
	fmt.Printf("[Part 1] Valid passports: %d\n", len(validPassportsPart1))

	var validPassportsPart2 = validatePassports(passports, true)
	fmt.Printf("[Part 2] Valid passports: %d\n", len(validPassportsPart2))

	var elapsed = time.Since(start)
	fmt.Printf("It took %s\n", elapsed)
}
