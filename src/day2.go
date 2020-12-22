package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	// advent helper functions
	"sojourner.me/advent2020/utils"
)

type Password struct {
	min, max int
	char string
	password string
}

func parseInput(input []string) []Password {
	var passwords []Password

	re, err := regexp.Compile(`(\d+)-(\d+) (\w): (\w+)`)
	if err != nil {
		println(err)
	}

	for _, line := range input {
		match := re.FindStringSubmatch(line)
		if match != nil {
			min, _ := strconv.Atoi(match[1])
			max, _ := strconv.Atoi(match[2])
			pw := Password{
				min,
				max,
				match[3],
				match[4],
			}
			passwords = append(passwords, pw)
		}
	}

	return passwords
}

func validPasswordPart1(pw Password) bool {
	count := strings.Count(pw.password, pw.char)
	return (count >= pw.min) && (count <= pw.max)
}

func validPasswordPart2(pw Password) bool {
	// min ^ max
	return (pw.password[pw.min - 1] == pw.char[0]) != (pw.password[pw.max - 1] == pw.char[0])
}

//gopls: ignore
func main() {
	lines := utils.ReadFile("input/day2")
	// lines := strings.Split(`
// 1-3 a: abcde
// 1-3 b: cdefg
// 2-9 c: ccccccccc
// `, "\n")
	items := parseInput(lines)
	count1, count2 := 0, 0
	for _, item := range items {
		if validPasswordPart1(item) {
			count1 += 1
		}
		if validPasswordPart2(item) {
			count2 += 1
		}
	}
	fmt.Printf("Part 1: %d, Part 2: %d\n", count1, count2)
}
