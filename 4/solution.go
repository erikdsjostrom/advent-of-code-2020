package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func contains(ss []string, s string) bool {
	for _, val := range ss {
		if val == s {
			return true
		}
	}
	return false
}

func clear_values__concat_and_sort(long_string string) string {
	//usr:asd etc:fs -> usretc
	var short_string string
	for _, s := range strings.Split(long_string, " ") {
		short_string += strings.Split(s, ":")[0]
	}
	return SortString(short_string)
}

func part1() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	correct := 0
	var line string
	for scanner.Scan() {
		if scanner.Text() == "" {
			pass := clear_values__concat_and_sort(line)
			if pass == "bcccddeeghhiiillprrrtyyy" || pass == "bccdeeghhiillprrrtyyy" {
				correct++
			}
			line = ""
		} else {
			line += " " + scanner.Text()
		}
	}
	fmt.Printf("Correct passports: %d\n", correct)
}

var h_color, _ = regexp.Compile("[a-f0-9]{6}")
var e_color = [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
var pid, _ = regexp.Compile("[0-9]{9}")

type comparable func(s string) bool

var is_valid = map[string]comparable{
	"byr": func(s string) bool {
		i, _ := strconv.Atoi(s)
		return i >= 1920 && i <= 2002
	},
	"iyr": func(s string) bool {
		i, _ := strconv.Atoi(s)
		return i >= 2010 && i <= 2020
	},
	"eyr": func(s string) bool {
		i, _ := strconv.Atoi(s)
		return i >= 2020 && i <= 2030
	},
	"hgt": func(s string) bool {
		height, _ := strconv.Atoi(s[:len(s)-2])
		measurement := s[len(s)-2:]
		switch measurement {
		case "cm":
			return height >= 150 && height <= 193 // skitsnack
		case "in":
			return height >= 59 && height <= 76
		default:
			fmt.Println(s)
			return false
		}
	},
	"hcl": func(s string) bool {
		return s[0] == '#' && h_color.MatchString(s[1:]) && len(s) == 7
	},
	"ecl": func(s string) bool {
		for _, ec := range e_color {
			if s == ec {
				return true
			}
		}
		return false
	},
	"pid": func(s string) bool {
		return pid.MatchString(s) && len(s) == 9
	},
	"cid": func(s string) bool {
		return true
	},
}

func part1_validation(pass string) bool {
	sorted_fields := clear_values__concat_and_sort(pass)
	return sorted_fields == "bcccddeeghhiiillprrrtyyy" || sorted_fields == "bccdeeghhiillprrrtyyy"
}

func part2_validation(pass string) bool {
	for _, field := range strings.Split(pass, " ") {
		pair := strings.Split(field, ":")
		if !is_valid[pair[0]](pair[1]) {
			return false
		}
	}
	return true
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	var pass string
	var part1 int
	var part2 int
	for scanner.Scan() {
		if scanner.Text() == "" {
			pass = strings.TrimSpace(pass)
			if part1_validation(pass) {
				part1++
				if part2_validation(pass) {
					part2++
				}
			}
			// Reset
			pass = ""
		} else {
			pass += " " + scanner.Text()
		}
	}
	pass = strings.TrimSpace(pass)
	if part1_validation(pass) {
		part1++
		if part2_validation(pass) {
			part2++
		}
	}
	// Reset
	pass = ""
	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", part2)
}
