package main

import (
	"ca/ca/elementary-ca/rules"
	"fmt"
	"time"
)

func main() {
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
	time.Sleep(time.Millisecond * 50)

	// T := 20

	for {
		world = rules.Rule110(currentWorld, world, n)
		fmt.Println(world)
		time.Sleep(time.Millisecond * 50)
	}
}
