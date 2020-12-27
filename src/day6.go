package main

import (
	"fmt"
	"strings"
	"github.com/deckarep/golang-set"
	
	"sojourner.me/advent2020/utils"
)

func main() {
	groups := strings.Split(utils.SlurpFile("input/day6"), "\n\n")

	// groups = strings.Split(`abc

// a
// b
// c

// ab
// ac

// a
// a
// a
// a

// b
// ba`, "\n\n")

	// part 1
	var total int
	for _, group := range groups {
		answers := mapset.NewSet()
		for _, char := range strings.ReplaceAll(group, "\n", "") {
			answers.Add(char)
		}
		total += answers.Cardinality()
	}
	fmt.Printf("Part 2: total count: %d\n", total)

	// part 2
	var total2 int
	for _, group := range groups {
		// create a set of each person's answers
		answers := strings.Split(group, "\n")
		answerSets := make([]mapset.Set, len(answers))
		for i, answer := range answers {
			answerSets[i] = mapset.NewSet()
			// add each answerer's answers to a set
			for _, char := range answer { answerSets[i].Add(char) }
		}
		// compute intersection of all the elements for this group
		totalSet := answerSets[0]
		for _, set := range answerSets {
			totalSet = totalSet.Intersect(set)
		}
		total2 += totalSet.Cardinality()
	}
	fmt.Printf("Part 2: total count: %d\n", total2)
}
