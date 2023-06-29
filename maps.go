package main

import "fmt"

func main() {
	// map1 := map[string]int{
	// 	"a": 1,
	// 	"b": 2,
	// 	"c": 3,
	// }
	// for k, v := range map1 {
	// 	fmt.Println(k, v)
	// }
	// fmt.Println(map1["a"])
	map1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	var map2 map[int]int
	map2[1] = 2
	fmt.Println(map2)
}
