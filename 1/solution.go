package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Why not over complicate?
func binary_search(ints []int, to_find int) (bool, int) {
	mid := len(ints) / 2
	switch {
	case len(ints) == 0:
		return false, -1
	case ints[mid] > to_find:
		return binary_search(ints[:mid], to_find)
	case ints[mid] < to_find:
		return binary_search(ints[mid+1:], to_find)
	default:
		return true, mid
	}
}

func main() {
	entries := []int{}
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		entries = append(entries, num)
	}
	sort.Ints(entries)

	// Part 1
	for current_index, value := range entries {
		to_find := 2020 - value
		found, found_index := binary_search(entries, to_find)
		if found && found_index != current_index {
			fmt.Printf("Part 1: %d\n", value*to_find)
			break
		}
	}

	// Part 2
	for i1, v1 := range entries {
		for i2, v2 := range entries {
			if v1+v2 > 2020 || i1 == i2 {
				break
			}
			to_find := 2020 - v1 - v2
			found, found_index := binary_search(entries, to_find)
			if found && found_index != i1 && found_index != i2 {
				fmt.Printf("Part 2: %d\n", v1*v2*to_find)
				break
			}
		}
	}
}
