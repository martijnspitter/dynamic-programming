package main

import (
	"fmt"
	"time"
)

func main() {
	checkWithLength()
	checkMemoWithLength()
	// checkMemo()
}

func checkWithLength() {
	defer timer(time.Now(), "withLength")
	v, err := bruteForceBestSumWithLengthCheck(49, []int{3, 4, 5, 7})
	if err != nil {
		fmt.Printf("%s ", err)
	} else {
		fmt.Printf("%v ", v)
	}
}

func checkMemoWithLength() {
	defer timer(time.Now(), "memo")
	summer := bestSum{
		memo: map[int][]int{},
	}

	v, err := summer.memoizedBestSumWithLengthCheck(49, []int{3, 4, 5, 7})
	if err != nil {
		fmt.Printf("%s ", err)
	} else {
		fmt.Printf("%v ", v)
	}
}

// func checkMemo() {
// 	defer timer(time.Now(), "memo COPILOT")
// 	summer := bestSum{
// 		memo: map[int][]int{},
// 	}
//
// 	v := summer.memoizedBestSum(49, []int{3, 4, 5, 7})
// 	fmt.Printf("%v COPILOT", v)
// }

func bruteForceBestSumWithLengthCheck(n int, nums []int) ([]int, error) {
	if n == 0 {
		return []int{}, nil
	}

	if n < 0 {
		return []int{}, fmt.Errorf("Can't sum this mate")
	}

	shortestSolution := []int{}

	for _, v := range nums {
		remainder := n - v
		res, err := bruteForceBestSumWithLengthCheck(remainder, nums)
		if err == nil {
			// check length
			partialSolution := append(res, v)
			if len(shortestSolution) == 0 || len(partialSolution) < len(shortestSolution) {
				shortestSolution = partialSolution
			}
		}
	}

	if len(shortestSolution) == 0 {
		return shortestSolution, fmt.Errorf("Can't sum this mate")
	}

	return shortestSolution, nil
}

func timer(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}

type bestSum struct {
	memo map[int][]int
}

func (bestSum bestSum) memoizedBestSumWithLengthCheck(n int, nums []int) ([]int, error) {
	v, ok := bestSum.memo[n]
	if ok {
		return v, nil
	}
	if n == 0 {
		return []int{}, nil
	}

	if n < 0 {
		return []int{}, fmt.Errorf("Can't sum this mate")
	}

	var shortestSolution []int = nil

	for _, v := range nums {
		remainder := n - v
		res, err := bestSum.memoizedBestSumWithLengthCheck(remainder, nums)
		if err == nil {
			// check length
			partialSolution := append(res, v)
			if shortestSolution == nil || len(partialSolution) < len(shortestSolution) {

				fmt.Println("Partial solution n ", n, partialSolution)
				fmt.Println("Shortest solution n ", n, shortestSolution)
				shortestSolution = partialSolution
			}
		}
	}
	fmt.Println("Saving shortestSolution n ", n, shortestSolution)
	bestSum.memo[n] = shortestSolution

	if shortestSolution == nil {
		return shortestSolution, fmt.Errorf("Can't sum this mate")
	}

	return shortestSolution, nil
}
