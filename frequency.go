package main

import "fmt"

func findfreq(arr []int) {
	freq := make(map[int]int)
	fmt.Println(freq)
	for _, item := range arr {
		if freq[item] == 0 {
			freq[item] = 1
		} else {
			freq[item]++
		}
	}
	for item, count := range freq {
		fmt.Println(item, "->", count)
	}

}

func main() {
	findfreq([]int{1, 1, 2, 3})
}
