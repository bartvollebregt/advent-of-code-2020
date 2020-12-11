package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func readFileWithReadString(filename string) (lines []int) {

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
			number, _ := strconv.Atoi(line)
			lines = append(lines, number)
		}

		if err != nil {
			break
		}
	}

	return lines
}

func main() {
	start := time.Now()

	var numbers = readFileWithReadString("./day_10/input.txt")
	var max = arrayMax(numbers)
	numbers = append(numbers, max+3)
	var joltDifference = findJoltDifference(numbers)
	fmt.Println(joltDifference)
	var distinctiveArrangements = findDistinctiveArrangements(numbers)
	fmt.Println(distinctiveArrangements)

	var elapsed = time.Since(start)
	fmt.Printf("It took %s\n", elapsed)
}

func findDistinctiveArrangements(numbers []int) int {
	sort.Ints(numbers)
	sol := map[int]int{0: 1}

	for _, number := range numbers {
		sol[number] = 0
		if _, ok := sol[number-1]; ok {
			sol[number] += sol[number-1]
		}
		if _, ok := sol[number-2]; ok {
			sol[number] += sol[number-2]
		}
		if _, ok := sol[number-3]; ok {
			sol[number] += sol[number-3]
		}
	}

	return sol[numbers[len(numbers)-1]]
}

func findJoltDifference(numbers []int) (difference int) {
	sort.Ints(numbers)
	sol := map[int]int{1: 0, 2: 0, 3: 0}
	prev := 0
	for _, number := range numbers {
		sol[number-prev] += 1
		prev = number
	}
	return sol[1] * sol[3]
}

func arrayMax(array []int) int {
	max := array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}
