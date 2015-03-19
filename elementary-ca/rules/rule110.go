package rules

func Rule110(currentWorld, world []int, n int) []int {
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

	return world
}
