package main

import "fmt"

// GridTraveler calculates the number of ways to travel a grid of size m x n

func main() {
	traveler := GridTraveler{memo: make(map[string]int)}

	fmt.Println(traveler.calculatePossiblePaths(2, 3))
	fmt.Println(traveler.calculatePossiblePaths(3, 3))
	fmt.Println(traveler.calculatePossiblePaths(4, 4))
	fmt.Println(traveler.calculatePossiblePaths(1000, 1000))
}

type GridTraveler struct {
	memo map[string]int
}

func (gt GridTraveler) calculatePossiblePaths(x int, y int) int {
	keypart := fmt.Sprintf("%f", x)
	keySecondPart := fmt.Sprintf("%f", y)

	key := keypart + "," + keySecondPart
	val, ok := gt.memo[key]

	if ok {
		return val
	}

	if x == 0 || y == 0 {
		return 0
	}
	if x == 1 && y == 1 {
		return 1
	}
	res := gt.calculatePossiblePaths(x-1, y) + gt.calculatePossiblePaths(x, y-1)

	gt.memo[key] = res

	return res
}
