package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFileWithReadString(filename string) (lines []string) {

	file, err := os.Open(filename)
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}

		if err != nil {
			break
		}
	}

	return lines
}

func getValidPasswords(passwords []string) (validPart1Passwords []string, validPart2Passwords []string) {
	var format = regexp.MustCompile(`(\d+)-(\d+)\s+(\w):\s+([\w\d]+)`)
	for _, password := range passwords {
		matches := format.FindStringSubmatch(password)
		least, _ := strconv.Atoi(matches[1])
		most, _ := strconv.Atoi(matches[2])
		character := matches[3]
		password := matches[4]

		var passwordIsValidPart1 = matchesPart1(least, most, character, password)
		var passwordIsValidPart2 = matchesPart2(least, most, character, password)

		if passwordIsValidPart1 {
			validPart1Passwords = append(validPart1Passwords, password)
		}

		if passwordIsValidPart2 {
			validPart2Passwords = append(validPart2Passwords, password)
		}

	}
	return validPart1Passwords, validPart2Passwords
}

func matchesPart1(least int, most int, character string, password string) (matches bool) {
	var passwordR = regexp.MustCompile(character)
	var passwordMatchCount = len(passwordR.FindAllStringSubmatch(password, -1))
	return passwordMatchCount >= least && passwordMatchCount <= most
}

func matchesPart2(least int, most int, character string, password string) (matches bool) {
	var leastCharacter = string([]rune(password)[least-1])
	var mostCharacter = string([]rune(password)[most-1])
	if leastCharacter == character && mostCharacter == character {
		return false
	}

	if leastCharacter != character && mostCharacter != character {
		return false
	}

	return true
}

func main() {
	var lines = readFileWithReadString("./day_2/input.txt")
	var matchesPart1, matchesPart2 = getValidPasswords(lines)
	fmt.Printf("We found %d Part 1 matches\n", len(matchesPart1))
	fmt.Printf("We found %d Part 2 matches\n", len(matchesPart2))
}
