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

func traverseTree(tree map[string][]Bag, seen *mapset.Set, key string, count *int) bool {
	// go through each sub-branch on this node
	retval := false
	for _, branch := range tree[key] {
		// traverse to the sub-branch
		parentSeen := traverseTree(tree, seen, branch.name, count)
		if branch.name == "shiny gold" || parentSeen {
			if !(*seen).Contains(key) {
				(*seen).Add(key)
				*count++
			}
			retval = true
		}
	}
	return retval
}

func traverseRoot(tree map[string][]Bag, root string) int {
	if len(tree[root]) == 0 {
		// leaf, so there's just one bag - this one
		return 1
	}
	var localCount int
	// recurse through the subtree
	for _, child := range tree[root] {
		// make sure to count the outer bag
		if len(tree[child.name]) > 0 {
			localCount += child.count
		}
		// as well as any bags contained within it
		ret := traverseRoot(tree, child.name)
		localCount += ret * child.count
	}
	return localCount
}

func main() {
	lines := utils.ReadFile("input/day7")

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
	// fmt.Printf("%v\n", tree)
	// part 1: traverse the tree
	var total int
	seen := mapset.NewSet()
	for _, k := range keys {
		// check decendents of every branch of the tree to find shiny gold
		traverseTree(tree, &seen, k, &total)
	}
	fmt.Printf("Part 1: %d\n", total)

	// part 2
	total2 := traverseRoot(tree, "shiny gold")
	fmt.Printf("Part 2: %d\n", total2)
}
