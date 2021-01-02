package main

import (
	"fmt"
	"sort"
	// "strings"

	"sojourner.me/advent2020/utils"
)

func main() {
	input := utils.ReadFile("input/day10")

	// turn input into integers and sort it
	int_input := utils.StringSliceToIntSlice(input)
	// prepend a zero rating for the charging outlet
	int_input = append([]int{0}, int_input...)
	sort.Ints(int_input)
	// calculate device's built-in joltage adapter rating, which is 3 more than max
	device_adapter := int_input[len(int_input)-1] + 3
	// add the built-in adapter to the end of the input
	int_input = append(int_input, device_adapter)

	counts := map[int]int {
		1: 0,
		2: 0,
		3: 0,
	}
	n := len(int_input)
	for i := range int_input {
		if i > n - 2 { break }
		// calculate diff and add it to the count map
		diff := int_input[i+1] - int_input[i]
		counts[diff]++
	}
	// part 1: (count of 1-jolt diff * count of 3-jolt diff)
	fmt.Println("Part 1:", counts[1] * counts[3])

	// counts2 is an array that tracks each successive count as we backtrack from the end
	counts2 := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		// 2nd to last and last have only 1 way to go
		if i > n - 3 {
			counts2[i] = 1
			continue
		}
		// check how many ways to go from this point
		total_dist := 0
		counts := 0
		for j := i+1; j < i + 4; j++ {
			delta := int_input[j] - int_input[j-1]
			total_dist += delta
			if total_dist > 3 {
				// if the distance to the next adapter is more than 3, stop counting
				break
			}
			counts++
		}
		// if there are more than one way to go, add up all of the ways
		if counts > 1 {
			new_count := 0
			for j := 0; j <= counts; j++ {
				new_count += counts2[i+j]
			}
			counts2[i] = new_count
		// otherwise, just use the previous entry
		} else {
			counts2[i] = counts2[i+1]
		}
	}
	fmt.Println("Part 2:", counts2[0])

}
