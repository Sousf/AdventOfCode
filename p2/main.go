package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		ranges := strings.Split(line, ",")
		for _, r := range ranges {
			input := strings.Split(r, "-")
			sum += sumInvalid(input)
		}
	}
	fmt.Printf("Sum invalid is: %d\n", sum)
}

func sumInvalid(s []string) int {
	upper, _ := strconv.Atoi(s[1])
	lower, _ := strconv.Atoi(s[0])
	currSum := 0
	for i := lower; i < upper; i++ {
		if isInvalidP2(strconv.Itoa(i)) {
			// sum up
			fmt.Println(i)
			currSum += i
		}
	}
	return currSum
}

func isInvalid(num string) bool {
	// if num starts with 0
	if num[0] == '0' {
		return true
	}
	// if num is a repeated pattern twice (has to have even length)
	if len(num)%2 != 0 {
		return false
	}
	middleIndex := len(num) / 2
	firstHalf, _ := strconv.Atoi(num[:middleIndex])
	secondHalf, _ := strconv.Atoi(num[middleIndex:])

	if firstHalf == secondHalf {
		return true
	}
	return false
}

func isInvalidP2(num string) bool {
	n := len(num)

	// Find the smallest length repeating pattern
	for patternLength := 1; patternLength <= n/2; patternLength++ {
		if n%patternLength != 0 {
			// divides not evenly, this pattern length is not possible
			continue
		}
		pattern := num[:patternLength]
		repeats := n / patternLength
		if repeats < 2 {
			continue
		}
		temp := ""
		for i := 0; i < repeats; i++ {
			temp = strings.Join([]string{temp, pattern}, "")
		}
		if temp == num {
			// found invalid
			return true
		}
	}
	return false
}
