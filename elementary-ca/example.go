package main

import (
	"ca/ca/elementary-ca/rules"
	"fmt"
)

func main() {
	// var center, left, right int
	n := 20
	world := make([]int, n)
	currentWorld := make([]int, n)

	// set initial state
	world[3] = 1
	world[4] = 1
	world[6] = 1
	world[7] = 1
	world[9] = 1
	world[10] = 1
	world[11] = 1
	fmt.Println(world)

	T := 20

	for j := 0; j < T; j++ {
		world = rules.Rule110(currentWorld, world, n)
		fmt.Println(world)
	}

}
