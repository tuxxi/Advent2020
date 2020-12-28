package main

import (
	"errors"
	"fmt"
	"github.com/deckarep/golang-set"
	"strconv"
	"strings"

	"sojourner.me/advent2020/utils"
)

func runVm(program []string) (int64, error) {
	var accumulator int64
	seen := mapset.NewSet()
	for i := 0; i < len(program); i++ {
		if len(program[i]) < 2 {
			continue
		}

		if seen.Contains(i) {
			return accumulator, errors.New("did not complete")
		}

		seen.Add(i)
		line := program[i]
		fields := strings.Fields(line)
		instr := fields[0]
		value, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			panic(err)
		}

		// run vm
		switch instr {
		case "acc":
			accumulator += value
		case "jmp":
			i += int(value) - 1
		case "nop":
			// nothing
		}
	}
	return accumulator, nil
}

func partOne(program []string) int64 {
	accum, _ := runVm(program)
	return accum
}
func partTwo(program []string) int64 {
	progCopy := make([]string, len(program))
	for i, line := range program {
		// switch nop to jmp or jmp to nop
		if len(line) < 2 {
			continue
		}

		// re-init progCopy
		copy(progCopy, program)
		fields := strings.Fields(line)

		switch fields[0] {
		case "acc":
			// do nothing
		case "jmp":
			progCopy[i] = "nop 0"
		case "nop":
			// replace this jump with
			progCopy[i] = "jmp " + fields[1]
		}
		accum, err := runVm(progCopy)
		if err != nil {
			// did not terminate, try again
			continue
		} else {
			return accum
		}
	}
	panic("None of the program variations terminated!")
}

func main() {
	program := utils.ReadFile("input/day8")

	fmt.Println("Part 1", partOne(program))
	fmt.Println("Part 2", partTwo(program))
}
