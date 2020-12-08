package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
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

	var lines = readFileWithReadString("./day_7/input.txt")
	var outerBags = mapToOuterBags(lines)
	var containingOuterBags = getOuterBagsContainingColor(outerBags, BAG_COLOR)

	fmt.Printf("[Part 1] Amount of bags containing '%s' bags: %d\n", BAG_COLOR, len(containingOuterBags))

	var requiredInnerBagAmount = getInnerBagAmount(outerBags, outerBags[BAG_COLOR]) - 1

	fmt.Printf("[Part 2] Amount of required inner Bags for bag '%s': %d\n", BAG_COLOR, requiredInnerBagAmount)

	var elapsed = time.Since(start)
	fmt.Printf("It took %s\n", elapsed)
}
func bagsContainsColor(outerBags OuterBags, bags []Bag, color string) bool {
	var containsColor = false
	for _, bag := range bags {
		if bag.color == color {
			return true
		} else if _, ok := outerBags[bag.color]; ok {
			if bagsContainsColor(outerBags, outerBags[bag.color], color) {
				containsColor = true
			}
		}
	}
	return containsColor
}

func getOuterBagsContainingColor(outerBags OuterBags, color string) (containingOuterBags []string) {
	for outerBagColor, innerBags := range outerBags {
		if bagsContainsColor(outerBags, innerBags, color) {
			containingOuterBags = append(containingOuterBags, outerBagColor)
		}
	}
	return containingOuterBags
}

func getInnerBagAmount(outerBags OuterBags, bags []Bag) (requiredInnerBagCount int) {
	requiredInnerBagCount = 1

	for _, bag := range bags {
		newInnerBags := outerBags[bag.color]
		bagAmount := bag.amount
		var childBagAmount = getInnerBagAmount(outerBags, newInnerBags)
		requiredInnerBagCount += childBagAmount * bagAmount
	}

	return requiredInnerBagCount
}

func mapToOuterBags(lines []string) (outerBags OuterBags) {
	outerBags = make(OuterBags)
	var bagRegex = regexp.MustCompile(`(\d+)\s+(.*?)\s+bag(?:s)?`)
	for _, line := range lines {
		var splitted = strings.Split(line, "bags contain")
		var outerBagColor = strings.TrimSpace(splitted[0])
		var innerBagsSplitted = strings.Split(strings.TrimSpace(splitted[1]), ",")
		var innerBags []Bag
		if strings.TrimSpace(splitted[1]) != "no other bags." {
			for _, bag := range innerBagsSplitted {
				var bagMatch = bagRegex.FindStringSubmatch(bag)
				var innerBagColor = bagMatch[2]
				var innerBagAmount, _ = strconv.Atoi(strings.TrimSpace(bagMatch[1]))
				var innerBag = Bag{
					color:  innerBagColor,
					amount: innerBagAmount,
				}
				innerBags = append(innerBags, innerBag)
			}
		}
		outerBags[outerBagColor] = innerBags
	}
	return outerBags
}
