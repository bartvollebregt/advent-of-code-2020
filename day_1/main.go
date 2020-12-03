package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const SUM = 2020

func readFileWithReadString(filename string) (numbers []int) {

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
			var number, _ = strconv.Atoi(line)
			numbers = append(numbers, number)
		}

		if err != nil {
			break
		}
	}

	return numbers
}

func findTwoNumberSumProduct(numbers []int, sum int) (num1 int, num2 int, sumResult int, product int) {
	for index1, num1 := range numbers {
		for index2, num2 := range numbers {
			if index1 == index2 {
				continue
			}
			var sumResult = num1 + num2
			if sumResult == sum {
				var product = num1 * num2
				return num1, num2, sumResult, product
			}
		}
	}
	return
}

func findThreeNumberSumProduct(numbers []int, sum int) (num1 int, num2 int, num3 int, sumResult int, product int) {
	for index1, num1 := range numbers {
		for index2, num2 := range numbers {
			for index3, num3 := range numbers {
				if index1 == index2 || index2 == index3 || index1 == index3 {
					continue
				}
				var sumResult = num1 + num2 + num3
				if sumResult == sum {
					var product = num1 * num2 * num3
					return num1, num2, num3, sumResult, product
				}
			}
		}
	}
	return
}

func main() {
	var lines = readFileWithReadString("./day_1/input.txt")
	{
		var num1, num2, sumResult, product = findTwoNumberSumProduct(lines, SUM)
		fmt.Printf("[Part1] We've got a winner! %d + %d = %d, product = %d\n", num1, num2, sumResult, product)
	}
	{
		var num1, num2, num3, sumResult, product = findThreeNumberSumProduct(lines, SUM)
		fmt.Printf("[Part2] We've got a winner! %d + %d + %d = %d, product = %d\n", num1, num2, num3, sumResult, product)
	}
}
