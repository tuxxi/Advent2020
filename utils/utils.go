package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filename string) []string {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	var retval []string
	for scanner.Scan() {
		retval = append(retval, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return retval
}