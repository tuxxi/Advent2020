package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	// advent helper functions
	"sojourner.me/advent2020/utils"
)

var boolToInt = map[bool]int{false: 0, true: 1}

func numInRange(s string, l, h int) bool {
	n, err := strconv.Atoi(s)
	return err == nil && l <= n && n <= h
}

var valid_map = map[string]func(s string) bool{
	"byr": func(s string) bool { return len(s) == 4 && numInRange(s, 1920, 2002) },
	"iyr": func(s string) bool { return len(s) == 4 && numInRange(s, 2010, 2020) },
	"eyr": func(s string) bool { return len(s) == 4 && numInRange(s, 2020, 2030) },
	"hgt": func(s string) bool {
		// smallest valid string is 4 chars long, like '60in'.
		if len(s) < 4 {
			return false
		}
		value, unit := s[:len(s)-2], s[len(s)-2:]
		if unit == "in" && numInRange(value, 59, 76) {
			return true
		}
		if unit == "cm" && numInRange(value, 150, 193) {
			return true
		}
		return false
	},
	"hcl": func(s string) bool {
		_, err := hex.DecodeString(strings.TrimPrefix(s, "#"))
		return err == nil && len(s) == 7
	},
	"ecl": func(s string) bool {
		switch s {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			return true
		}
		return false
	},
	"pid": func(s string) bool {
		_, err := strconv.ParseFloat(s, 64)
		return err == nil && len(s) == 9
	},
}

func main() {
	// we want to read the entire file into a string
	file := utils.SlurpFile("input/day4")
	lines := strings.Split(string(file), "\n\n")
	fmt.Printf("There are %d total passports\n\n", len(lines))

	var cnt1, cnt2 int
	for _, itm := range lines {
		var valid1, valid2 int
		for _, field := range strings.Fields(itm) {
			kv := strings.Split(field, ":")
			// check is key is valid
			valid, ok := valid_map[kv[0]]
			// check if key exists in map for part 1
			if ok {
				valid1++
				// check if validator function returns a valid result for part 2
				valid2 += boolToInt[valid(kv[1])]
			}
		}
		cnt1 += boolToInt[valid1 >= 7]
		cnt2 += boolToInt[valid2 >= 7]
	}

	fmt.Println("part 1:", cnt1, "valid passports")
	fmt.Println("part 2:", cnt2, "valid passports")

}
