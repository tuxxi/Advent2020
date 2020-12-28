package main

import (
	"fmt"
	"sort"

	"sojourner.me/advent2020/utils"
)

func checkSum(slice []int, target int) bool {
	for i := range slice {
		for j := range slice {
			if slice[i]+slice[j] == target {
				return true
			}
		}
	}
	return false
}

func main() {
	input := utils.ReadFile("input/day9")

	preambleSize := 25
	var invalidNum int
	for i := range input {
		var prevSlice []int
		if i+preambleSize < len(input) {
			prevSlice = utils.StringSliceToIntSlice(input[i : i+preambleSize])
			num := utils.ParseInt(input[i+preambleSize])
			if !checkSum(prevSlice, num) {
				invalidNum = num
				break
			}
		}
	}
	fmt.Println("Part 1:", invalidNum)

	// part 2
	var slice []int
	n := len(input)
	out: for i := 0; i < n; i++ {
		for j := i+2; j < n; j++ {
			slice = utils.StringSliceToIntSlice(input[i:j])
			sum := utils.SumOfSlice(slice)
			// fmt.Println(i, j, slice, sum)
			if sum == invalidNum {
				break out
			}
		}
	}

	// find sum of smallest and largest number in the slice
	sort.Ints(slice)
	sum := slice[0] + slice[len(slice)-1]
	fmt.Println("Part 2:", sum)
}
