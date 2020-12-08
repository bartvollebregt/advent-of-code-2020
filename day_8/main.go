package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const BAG_COLOR = "shiny gold"

type OuterBags = map[string][]Bag

type Bag = struct {
	color  string
	amount int
}

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

func main() {
	start := time.Now()

	var lines = readFileWithReadString("./day_8/input.txt")
	var acc1, _ = executeBootCode(lines)
	fmt.Printf("[Part 1] Acc after stopping before second row execution: %d\n", acc1)

	var acc2 = findAndFixLoop(lines)

	fmt.Printf("[Part 2] Acc after fixing code: %d\n", acc2)

	var elapsed = time.Since(start)
	fmt.Printf("It took %s\n", elapsed)
}

func findAndFixLoop(lines []string) (acc int) {
	for index, line := range lines {
		operation, argument := parseCodeLine(line)
		if operation == "acc" {
			continue
		}
		newCode := make([]string, len(lines))
		copy(newCode, lines)

		switch operation {
		case "nop":
			newCode[index] = fmt.Sprintf("jmp %+d", argument)
		case "jmp":
			newCode[index] = fmt.Sprintf("nop %+d", argument)
		default:
			fmt.Printf("Unknown operation %s\n", operation)
			os.Exit(-1)
		}

		acc, stopped := executeBootCode(newCode)
		if !stopped {
			return acc
		}
	}

	return 0
}

func executeBootCode(lines []string) (acc int, stopped bool) {
	executedLines := make(map[int]bool)
	for i := 0; i < len(lines); i++ {
		if executedLines[i] {
			return acc, true
		}
		operation, argument := parseCodeLine(lines[i])
		executedLines[i] = true
		switch operation {
		case "acc":
			acc += argument
		case "nop":
			continue
		case "jmp":
			i += argument - 1
		default:
			fmt.Printf("Unknown operation %s\n", operation)
			os.Exit(-1)
		}
	}

	return acc, false
}

func parseCodeLine(line string) (operation string, argument int) {
	var splitLine = strings.Split(line, " ")
	operation = splitLine[0]
	argument, _ = strconv.Atoi(splitLine[1])
	return operation, argument
}
