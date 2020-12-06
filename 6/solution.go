package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	var yes_answers = map[rune]int{}
	var sum_part1 int
	var sum_part2 int
	var group_members int
	for scanner.Scan(){
		line := scanner.Text()
		if line == "" {
			sum_part1 += len(yes_answers)
			for _, number_of_answers := range yes_answers {
				if number_of_answers == group_members {
					sum_part2++
				}
			}
			// Reset
			group_members = 0
			yes_answers = map[rune]int{}
		} else {
			group_members++
			for _, c := range line {
				yes_answers[c]++
			}
		}
	}
	sum_part1 += len(yes_answers)
	for _, number_of_answers := range yes_answers {
		if number_of_answers == group_members {
			sum_part2++
		}
	}
	fmt.Printf("Part 1: %d\n", sum_part1)
	fmt.Printf("Part 2: %d\n", sum_part2)
}
