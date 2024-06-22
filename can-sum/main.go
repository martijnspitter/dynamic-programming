package main

import "fmt"

func main() {
	sum := sum{memo: make(map[int]bool)}
	fmt.Printf("%t", sum.canSum(30000000, []int{7, 15, 21, 28}))
}

type sum struct {
	memo map[int]bool
}

func (sum sum) canSum(n int, nums []int) bool {
	val, ok := sum.memo[n]

	if ok {
		return val
	}

	if n == 0 {
		return true
	}
	if n < 0 {
		return false
	}

	for i := 0; i < len(nums); i++ {
		remainder := n - nums[i]
		if sum.canSum(remainder, nums) {
			sum.memo[n] = true

			return true
		}
	}

	sum.memo[n] = false
	return false
}
