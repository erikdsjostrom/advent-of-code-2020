package main

import (
	"bufio"
	"fmt"
	"os"
)

type slope struct {
	x int
	y int
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	line_length := 31
	trees := []int{0, 0, 0, 0, 0}
	slopes := []slope{
		{x: 1, y: 1},
		{x: 3, y: 1},
		{x: 5, y: 1},
		{x: 7, y: 1},
		{x: 1, y: 2},
	}
	tree := rune('#')
	row := 0
	for scanner.Scan() {
		line := []rune(scanner.Text())
		for i, slope := range slopes {
			if (row%slope.y == 0) && line[(slope.x*row)%line_length] == tree {
				trees[i] += 1
			}
		}
		row++
	}
	total_trees := 1
	for _, num := range trees {
		total_trees *= num
	}
	fmt.Printf("Part 1: %d\nPart 2: %d\n", trees[1], total_trees)
}
