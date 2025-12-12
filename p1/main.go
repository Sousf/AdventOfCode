package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	curr := 50
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		// do stuff with input
		curr = executeInstruction(line, curr)
		fmt.Printf("Instruction : %s, Result: %d \n", line, curr)
		if curr == 0 {
			count++
			fmt.Printf("Count: %s", strconv.Itoa(count))
		}
	}
	fmt.Printf("Secret Password is: %d \n", count)
}

func executeInstruction(instruction string, curr int) int {
	numStr := instruction[1:]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Printf("Error converting %v to int type: \n", numStr)
	}
	if instruction[0] == 'R' {
		curr += num
	} else {
		curr -= num
	}

	return ((curr % 100) + 100) % 100
}
