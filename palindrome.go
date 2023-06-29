package main

import (
	"fmt"
)

func isPalindrome(n int) bool {
	temp := n
	rev := 0
	for temp > 0 {
		rem := temp % 10
		rev = rev*10 + rem
		temp = temp / 10
	}
	if rev == n {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(132))
}
