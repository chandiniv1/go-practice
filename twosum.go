package main

import "fmt"

func twoSum(nums []int, target int) []int {
	n := len(nums)
	arr := []int{0, 0}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				arr[0] = i
				arr[1] = j
			}
		}
	}
	return arr
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	indices := twoSum(nums, target)
	fmt.Println(indices)
}
