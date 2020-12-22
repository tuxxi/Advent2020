package main

import (
	"fmt"
	// advent helper functions
	"sojourner.me/advent2020/utils"
)

func traverse(t Traversal, lines []string) int {
	// position in the grid
	x, y := 0, 0
	// number of trees seen so far
	trees := 0
	// size of the grid
	max_x, max_y := len(lines[0]), len(lines)
	for {
		if y >= max_y {
			break
		}
		line := lines[y]
		if line[x] == '#' {
			trees += 1
		}
		x = (x + t.x_step) % max_x
		y += t.y_step
	}
	return trees
}

type Traversal struct {
	x_step, y_step int
}
//gopls: ignore
func main() {
	lines := utils.ReadFile("input/day3")
	// part 1
	p1 := traverse(Traversal{3, 1}, lines)
	fmt.Println(p1)
	// part 2
	traversals := []Traversal{
		Traversal{1,1},
		Traversal{3,1}, 
		Traversal{5,1}, 
		Traversal{7,1}, 
		Traversal{1,2},
	}
	total := 1
	for _, t := range traversals {
		trees := traverse(t, lines)
		fmt.Printf("%+v: %d\n", t, trees)
		total *= trees
	}
	fmt.Println(total)

}
