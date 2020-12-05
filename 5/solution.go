package main

import (
	"bufio"
	"fmt"
	"os"
)

var binary_sequence [7]int = [7]int{
	64,
	32,
	16,
	8,
	4,
	2,
	1,
}

func from_airplane_binary(str string) (num int) {
	for i, c := range str {
		if c == 'B' || c == 'R' {
			// 7 - len(str) is there so sequences of different lengths
			// will start at the correct index in the array,
			// 1101 starts at 3, 101 starts at 4 and so on
			num += binary_sequence[i+7-len(str)]
		}
	}
	return
}

func calc_seat_ID(str string) int {
	return from_airplane_binary(str[:7])*8 + from_airplane_binary(str[7:])
}

const number_of_seats = 128 * 8

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	var seats = [number_of_seats]bool{}
	var highest int = 0
	var seat_id int
	for scanner.Scan() {
		seat_id = calc_seat_ID(scanner.Text())
		if seat_id > highest {
			highest = seat_id
		}
		seats[seat_id] = true
	}
	fmt.Printf("Part 1: %d\n", highest)
	for i, seat := range seats {
		if i == 0 || i == number_of_seats-1 {
			continue
		}
		if seats[i-1] && !seat && seats[i+1] {
			fmt.Printf("Part 2: %d\n", i)
			break
		}
	}
}
