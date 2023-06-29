package main

import "fmt"

func main() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	var a = [3]int{1, 2, 3}
	//fmt.Println("dsdfd",a)
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	var arr1 [2]int
	arr1[0] = 1
	arr1[1] = 2
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
}
