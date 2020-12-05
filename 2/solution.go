package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)



func main() {
        password_pattern, _ := regexp.Compile("(\\d+)-(\\d+) (\\w): (\\w*)")
        file, _ := os.Open("./input.txt")
        scanner := bufio.NewScanner(file)
        ok_passwords_part_1 := 0
        ok_passwords_part_2 := 0
        for scanner.Scan() {
                groups := password_pattern.FindStringSubmatch(scanner.Text())
                min, _ := strconv.Atoi(groups[1])
                max, _ := strconv.Atoi(groups[2])
                letter := rune(groups[3][0])
                password := []rune(groups[4])
                occurrences := 0
                for _, c :=  range password {
                        if c == letter {
                                occurrences++
                        }
                }
                if occurrences >= min && occurrences <= max {
                        ok_passwords_part_1++
                }

                if (letter == password[min-1]) != (letter == password[max-1]) {
                        ok_passwords_part_2++
                }
        }
        fmt.Printf("Part 1: %d\n", ok_passwords_part_1)
        fmt.Printf("Part 2: %d\n", ok_passwords_part_2)
}
