package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func readFileWithReadFile(filename string) (file string) {

	b, err := ioutil.ReadFile(filename) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

func uniqueNonEmptyElementsOf(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}

	return us

}

func parseAnyoneGroupAnswers(file string) (answers [][]string) {
	var groups = strings.Split(file, "\n\n")
	for _, group := range groups {
		var people = strings.Split(group, "\n")
		var groupAnswers []string
		for _, person := range people {
			var yesAnswers = strings.Split(person, "")
			groupAnswers = append(groupAnswers, yesAnswers...)
		}
		groupAnswers = uniqueNonEmptyElementsOf(groupAnswers)
		answers = append(answers, groupAnswers)
	}

	return answers
}

func parseEveryoneGroupAnswers(file string) (answers [][]string) {
	var groups = strings.Split(file, "\n\n")
	for _, group := range groups {
		var people = strings.Split(strings.TrimSpace(group), "\n")
		var groupAnswersCount = make(map[string]int)
		var groupAnswers []string
		for _, person := range people {
			var yesAnswers = strings.Split(person, "")
			for _, yesAnswer := range yesAnswers {
				if _, ok := groupAnswersCount[yesAnswer]; !ok {
					groupAnswersCount[yesAnswer] = 0
				}
				groupAnswersCount[yesAnswer]++
			}

		}
		for answer, count := range groupAnswersCount {
			if count == len(people) {
				groupAnswers = append(groupAnswers, answer)
			}
		}
		answers = append(answers, groupAnswers)
	}
	return answers
}

func answerSumCount(answers [][]string) (count int) {
	count = 0
	for _, group := range answers {
		count = count + len(group)
	}
	return count
}

func main() {
	start := time.Now()

	var file = readFileWithReadFile("./day_6/input_test.txt")
	{
		var groupAnswers = parseAnyoneGroupAnswers(file)
		var answerCount = answerSumCount(groupAnswers)
		fmt.Printf("[Part 1] The total unique answer count is: %d\n", answerCount)
	}

	{
		var groupAnswers = parseEveryoneGroupAnswers(file)
		var answerCount = answerSumCount(groupAnswers)
		fmt.Printf("[Part 2] The total unique answer count is: %d\n", answerCount)
	}

	var elapsed = time.Since(start)
	fmt.Printf("It took %s\n", elapsed)
}
