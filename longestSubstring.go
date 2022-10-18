package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var maxLnth, currLnth int // maximum length and current length
	swStart := 0 // sliding window start index
	runeLog := map[rune]int{} // map of all the symbols we've seen

	for ind, val := range(s) {
		if swEnd, ok := runeLog[val]; ok && swEnd >= swStart { // sliding window end index
			swStart = swEnd + 1 // moving the start index to the next character after previously seen
			currLnth = ind - swEnd // getting all the symbols between the same runes
		} else {
			currLnth += 1
		}

		if currLnth > maxLnth {
			maxLnth = currLnth
		}

		runeLog[val] = ind // rewriting with current index
	}

	return maxLnth
}

func main() {
	var str string
	fmt.Print("Enter the string: ")
	fmt.Scan(&str)
	fmt.Println("The longest substring is", lengthOfLongestSubstring(str))
}
