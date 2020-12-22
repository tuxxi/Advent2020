package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readFile(filename string) []int64 {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	var retval []int64
	for scanner.Scan() {
		v, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		retval = append(retval, v)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return retval
}

func main() {
	lines := readFile("input/day1")
	// part 1
	log.Println(" --> Part 1:")
	for i, x := range lines {
		for _, y := range lines[i:] {
			if x + y == 2020 {
				log.Println(x, y)
				log.Println(x*y)
			}
		}
	}
	log.Println(" --> Part 2:")
	for i, x := range lines {
		for j, y := range lines[i:] {
			for _, z := range lines[j:] {
				if x + y + z == 2020 {
					log.Println(x, y, z)
					log.Println(x*y*z)
				}
			}
		}
	}
}
