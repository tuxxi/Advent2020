package main

import (
	"fmt"
	"strings"
	"strconv"
	"sojourner.me/advent2020/utils"
)

func charToBin (r rune) rune {
	switch r {
	case 'F': return '0'
	case 'B': return '1'
	case 'L': return '0'
	case 'R': return '1'
	default:
		return 'A'
	}
}

func bspToInt(bsp string) (int, error) {
	resStr := strings.Map(charToBin, bsp)
	parsed, err := strconv.ParseInt(resStr, 2, 64)
	if err != nil { return 0, err }
	return int(parsed), nil
}

func seatId(str string) int {
	row, err := bspToInt(str[:7])
	if err != nil {
		fmt.Println(err)
		return 0
	}
	col, err := bspToInt(str[7:])
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return row * 8 + col
}

func testSeatId() {
	var test_cases = map[string]int {
		"FBFBBFFRLR": 357,
		"BFFFBBFRRR": 567,
		"FFFBBBFRRR": 119,
		"BBFFBBFRLL": 820,
		"BBBFBBFLLL": 944,
	}
	for k, v := range test_cases {
		res := seatId(k)
		if res != v {
			fmt.Printf("!! - FAIL - For '%s', got %d, expected %d\n", k, res, v)
		} else {
			fmt.Printf("Passed: '%s' => %d\n", k, res)
		}
	}
}

func main() {
	// testSeatId()
	lines := utils.ReadFile("input/day5")
	var min, max, sum int
	for _, line := range lines {
		seat_id := seatId(line)
		if min == 0 || seat_id < min {
			min = seat_id
		}
		if seat_id > max {
			max = seat_id
		}
		sum += seat_id
	}
	fmt.Printf("min is %d, max seat id is %d, sum is %d\n", min, max, sum)

	// part 2: finding the missing number
	// closed form sum of all numbers
	expected := (max * (max + 1)) / 2 - (min * (min - 1)) / 2
	// missing number is sum - expected
	missing := expected - sum
	fmt.Printf("missing seat ID is %d\n", missing)
}
