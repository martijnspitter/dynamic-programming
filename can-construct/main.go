package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// checkBruteForce()
	checkMemo()
}

func checkBruteForce() {
	defer timer(time.Now(), "brute force")
	res := bruteForce("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeet", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee", "eeeeeee"})
	fmt.Println("Can construct: ", res)
}

func checkMemo() {
	defer timer(time.Now(), "memo")
	cc := canConstruct{
		memo: map[string]bool{},
	}
	res := cc.construct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeet", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee", "eeeeeeet"})
	fmt.Println("Can construct: ", res)
}

func bruteForce(word string, wordbank []string) bool {
	if len(word) == 0 {
		return true
	}

	for _, v := range wordbank {
		index := strings.Index(word, v)
		if index == 0 {
			var subword string = ""
			subword = word[len(v):]

			if bruteForce(subword, wordbank) {
				return true
			}
		}
	}
	return false
}

func timer(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}

type canConstruct struct {
	memo map[string]bool
}

func (cc canConstruct) construct(word string, wordbank []string) bool {
	if v, ok := cc.memo[word]; ok {
		return v
	}

	if len(word) == 0 {
		return true
	}

	for _, v := range wordbank {
		index := strings.Index(word, v)
		if index == 0 {
			var subword string = ""
			subword = word[len(v):]

			res := cc.construct(subword, wordbank)
			if res {
				cc.memo[word] = res
				return res
			}
		}
	}
	cc.memo[word] = false
	return false
}
