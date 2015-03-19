package main

import (
	"fmt"
)

func main() {
	var center, left, right int
	n := 38
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

	T := 50

	for j := 0; j < T; j++ {
		copy(currentWorld, world)
		for i := 0; i < n; i++ {
			if i == 0 {
				left = currentWorld[n-1]
			} else {
				left = currentWorld[i-1]
			}
			if i == n-1 {
				right = currentWorld[0]
			} else {
				right = currentWorld[i+1]
			}
			center = currentWorld[i]

			switch {
			case left == 0 && center == 0 && right == 0:
				world[i] = 0
			case left == 0 && center == 0 && right == 1:
				world[i] = 1
			case left == 0 && center == 1 && right == 0:
				world[i] = 1
			case left == 0 && center == 1 && right == 1:
				world[i] = 1
			case left == 1 && center == 0 && right == 0:
				world[i] = 0
			case left == 1 && center == 0 && right == 1:
				world[i] = 1
			case left == 1 && center == 1 && right == 0:
				world[i] = 1
			case left == 1 && center == 1 && right == 1:
				world[i] = 0
			}
		}
		// fmt.Println(currentWorld)
		fmt.Println(world)
	}

}
