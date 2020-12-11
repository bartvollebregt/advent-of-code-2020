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

const PREAMBLE = 25
const PREVIOUS_NUMBERS = 25

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

	var numbers = readFileWithReadString("./day_9/input.txt")
	var invalidSum = findFirstInvalidSum(numbers)
	fmt.Printf("[Part 1] The invalid Sum is: %d\n", invalidSum)

	var contiguousSum = findContiguousSumNumbers(numbers, invalidSum)
	min, max := minMax(contiguousSum)
	fmt.Printf("[Part 2] The encryption weakness is: %d\n", min+max)

	var elapsed = time.Since(start)
	fmt.Printf("It took %s\n", elapsed)
}

func minMax(array []int) (int, int) {
	var max = array[0]
	var min = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func findContiguousSumNumbers(numbers []int, invalidSum int) (result []int) {
	for index1 := 0; index1 < len(numbers)-1; index1++ {
		for index2 := index1 + 1; index2 < len(numbers)-1; index2++ {
			stack := numbers[index1:index2]
			arraySum := sumArray(stack)
			if arraySum > invalidSum {
				break
			}
			if invalidSum == arraySum {
				return stack
			}
		}
	}
	return result
}

func sumArray(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func findFirstInvalidSum(numbers []int) (result int) {
	for i := PREAMBLE; i < len(numbers); i++ {
		sum := numbers[i]
		prevNumbers := numbers[i-PREVIOUS_NUMBERS : i]
		matches := false
		for _, number1 := range prevNumbers {
			for _, number2 := range prevNumbers {
				if (number1 + number2) == sum {
					matches = true
				}
			}
		}
		if !matches {
			return sum
		}
	}
	return 0
}
