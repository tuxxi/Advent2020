package utils

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
)

func SlurpFile(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func ReadFile(filename string) []string {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	var retval []string
	for scanner.Scan() {
		retval = append(retval, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return retval
}

func ParseInt(input string) int {
	val, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(val)
}

func StringSliceToIntSlice(strings []string) []int {
	out := make([]int, len(strings))
	for i, str := range strings {
		out[i] = ParseInt(str)
	}
	return out
}

func SumOfSlice(ints []int) int {
	var out int
	for _, item := range ints {
		out += item
	}
	return out
}
