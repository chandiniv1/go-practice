package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//greeter("world")
	websitelist := []string{
		"https://lco.dev",
		"https://google.com",
		"https://go.dev",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

// func greeter(s string){
// 	for i:=0;i<6;i++{
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("oops in endpoint")
	}
	fmt.Printf("%d status code is %s \n", res.StatusCode, endpoint)
}
