package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
	"github.com/deckarep/golang-set"

	"sojourner.me/advent2020/utils"
)

type Bag = struct {
	name  string
	count int
}

// parse a line into a map of bag name => bag contents
func parseBagRules(line string, outTree *map[string][]Bag) {
	if len(line) < 2 {
		return
	}
	color := strings.TrimSpace(line[:strings.Index(line, "bags")])
	contents := strings.Split(
		line[strings.Index(line, "contain")+len("contain"):],
		",",
	)
	for _, content := range contents {
		if strings.Contains(content, "no other bag") {
			continue
		}
		words := strings.Fields(content)
		// count is alwas the first 
		count, err := strconv.ParseInt(words[0], 10, 64)
		if err != nil {
			panic(err)
		}
		
		// name is always two words, the 2nd and 3rd fields
		name := strings.TrimSpace(strings.Join(words[1:3], " "))
		(*outTree)[color] = append((*outTree)[color], Bag{name, int(count)})
	}
}

func traverseTree(tree map[string][]Bag, seen *mapset.Set, key, target string, count *int) bool {
	// go through each sub-branch on this node
	retval := false
	for _, branch := range tree[key] {
		parentSeen := traverseTree(tree, seen, branch.name, target, count)
		if branch.name == target || parentSeen {
			if !(*seen).Contains(key) {
				fmt.Printf("'%s' bag contains %s. seen: %v\n", key, target, *seen)
				(*seen).Add(key)
				*count++
			}
			retval = true
		}
	}
	return retval
}

func main() {
	lines := utils.ReadFile("input/day7")
	// lines = strings.Split(`light red bags contain 1 bright white bag, 2 muted yellow bags.
// dark orange bags contain 3 bright white bags, 4 muted yellow bags.
// bright white bags contain 1 shiny gold bag.
// muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
// shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
// dark olive bags contain 3 faded blue bags, 4 dotted black bags.
// vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
// faded blue bags contain no other bags.
// dotted black bags contain no other bags.`, "\n")

	// build up a tree, with map from bag color => contents
	tree := make(map[string][]Bag)
	for _, line := range lines {
		parseBagRules(line, &tree)
	}
	// create a list of sorted keys in order
	keys := make([]string, len(tree))
	i := 0
	for k := range tree {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Printf("%v\n", tree)
	// part 1: traverse the tree
	var total int
	seen := mapset.NewSet()
	for _, k := range keys {
		// check decendents of every branch of the tree to find shiny gold
		traverseTree(tree, &seen, k, "shiny gold", &total)
	}

	fmt.Printf("%d\n", total)
}
