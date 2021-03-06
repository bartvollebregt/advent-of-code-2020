package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

const ROW_AMOUNT = 128
const COLUMN_AMOUNT = 8
const LOWER_HALF_CHAR = "F"
const UPPER_HALF_CHAR = "B"
const LEFT_HALF_CHAR = "L"
const RIGHT_HALF_CHAR = "R"

type boardingPass struct {
	row    int
	column int
	seatId int
}

func (p boardingPass) String() string {
	return fmt.Sprintf("{SeatId: %d, Row: %d, Column: %d}", p.seatId, p.row, p.column)
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

func parseBoardingPassNumber(code string, amount int, lowerHalfChar string, upperHalfChar string) (row int) {
	var lowerBound = 0
	var upperBound = amount
	for _, char := range code {
		var newValue = ((upperBound - lowerBound) / 2) + lowerBound
		if string(char) == upperHalfChar {
			lowerBound = newValue
		} else if string(char) == lowerHalfChar {
			upperBound = newValue
		}
	}
	return lowerBound
}

func calculateSeatId(row int, column int) (seatId int) {
	return (row * 8) + column
}

func parseBoardingPass(code string) (pass boardingPass) {
	var row = parseBoardingPassNumber(code[0:7], ROW_AMOUNT, LOWER_HALF_CHAR, UPPER_HALF_CHAR)
	var column = parseBoardingPassNumber(code[7:10], COLUMN_AMOUNT, LEFT_HALF_CHAR, RIGHT_HALF_CHAR)
	var seatId = calculateSeatId(row, column)
	pass = boardingPass{
		row:    row,
		column: column,
		seatId: seatId,
	}
	return pass
}

func parseBoardingPasses(lines []string) (boardingPasses []boardingPass) {
	for _, line := range lines {
		var boardingPass = parseBoardingPass(line)
		boardingPasses = append(boardingPasses, boardingPass)
	}
	return boardingPasses
}

func findMaxSeatId(boardingPasses []boardingPass) (max boardingPass) {
	max = boardingPasses[0]
	for _, boardingPass := range boardingPasses {
		if boardingPass.seatId > max.seatId {
			max = boardingPass
		}
	}
	return max
}

func findMissingSeatIdsPasses(boardingPasses []boardingPass) (missingSeatIds []int) {
	var prevSeatId = boardingPasses[0].seatId - 1
	for _, boardingPass := range boardingPasses {
		if (prevSeatId + 1) != boardingPass.seatId {
			missingSeatIds = append(missingSeatIds, boardingPass.seatId-1)
		}
		prevSeatId = boardingPass.seatId
	}
	return missingSeatIds
}

func sortBoardingPasses(boardingPasses []boardingPass) []boardingPass {
	sort.Slice(boardingPasses[:], func(i, j int) bool {
		return boardingPasses[i].seatId < boardingPasses[j].seatId
	})
	return boardingPasses
}

func main() {
	start := time.Now()

	var lines = readFileWithReadString("./day_5/input.txt")
	var boardingPasses = parseBoardingPasses(lines)
	var maxBoardingPass = findMaxSeatId(boardingPasses)

	fmt.Printf("[Part 1] Maximum SeatId is: %d\n", maxBoardingPass.seatId)

	boardingPasses = sortBoardingPasses(boardingPasses)
	var missingSeatIds = findMissingSeatIdsPasses(boardingPasses)
	fmt.Printf("[Part 2] Missing Seat ID's are: %d\n", missingSeatIds)

	var elapsed = time.Since(start)
	fmt.Printf("It took %s\n", elapsed)
}
