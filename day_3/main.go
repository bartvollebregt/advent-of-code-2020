package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const TREE_CHAR = "#"

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

func findTrees(mapLines []string, down int, right int) (count int) {
	count = 0
	var currentIndex = 0
	for index, line := range mapLines {
		if index%down != 0 {
			continue
		}
		if currentIndex > len(line)-1 {
			currentIndex = currentIndex - len(line)
		}
		var character = string([]rune(line)[currentIndex])

		currentIndex = currentIndex + right
		if character == TREE_CHAR {
			count++
		}
	}
	return count
}

func main() {
	start := time.Now()

	var mapLines = readFileWithReadString("./day_3/input.txt")

	var part1FoundTrees = findTrees(mapLines, 1, 3)

	var part21foundtrees = findTrees(mapLines, 1, 1)
	var part22foundtrees = findTrees(mapLines, 1, 3)
	var part23foundtrees = findTrees(mapLines, 1, 5)
	var part24foundtrees = findTrees(mapLines, 1, 7)
	var part25foundtrees = findTrees(mapLines, 2, 1)

	var part2Product = part21foundtrees * part22foundtrees * part23foundtrees * part24foundtrees * part25foundtrees

	fmt.Printf("[Part 1] Trees: %d\n", part1FoundTrees)
	fmt.Printf("[Part 2] Product: %d\n", part2Product)

	var elapsed = time.Since(start)
	fmt.Printf("It took %s\n", elapsed)
}
