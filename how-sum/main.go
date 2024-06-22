package main

import "fmt"

func main() {
	summer := sum{
		memo: map[int][]int{},
	}

	val, err := summer.canSum(70001, []int{4, 14})
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("%v", val)
	}
}

type sum struct {
	memo map[int][]int
}

func (sum sum) canSum(n int, nums []int) ([]int, error) {
	m, ok := sum.memo[n]
	if ok {
		if len(m) > 0 {
			return m, nil
		}
		return make([]int, 0), fmt.Errorf("Could not sum")
	}

	if n == 0 {
		return make([]int, 0), nil
	}

	if n < 0 {
		return make([]int, 0), fmt.Errorf("Could not sum")
	}

	for _, v := range nums {
		remainder := n - v
		res, err := sum.canSum(remainder, nums)
		if err == nil {
			sum.memo[n] = append(res, v)
			return sum.memo[n], nil
		}
	}

	sum.memo[n] = make([]int, 0)
	return make([]int, 0), fmt.Errorf("Could not sum")
}
