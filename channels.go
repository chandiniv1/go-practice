package main

import (
	"fmt"
	"sync"
)

func main() {
	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 6
		wg.Done()
	}(myCh, wg)
	wg.Wait()
	mych := make(chan int)
	go func() {
		mych <- 45
	}()
	val := <-mych
	fmt.Println(val)

	mychan1 := make(chan bool)
	go func() {
		mychan1 <- true
	}()
	fmt.Println(<-mychan1)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			//myChan2 <- 23
			// fmt.Println(<-myChan2) results in deadlock
			fmt.Println(i)
			defer wg.Done()
		}(i)
		wg.Wait()
	}

	wg.Add(2)
	myChan2 := make(chan int)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myChan2)   //it will print 5
		fmt.Println(<-myChan2)	//it will print 6 
		wg.Done()
	}(myChan2, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myChan2 <- 5
		myChan2 <- 6
		//fmt.Println(<-myChan2) if we keep the print statement here it results in deadlock
		wg.Done()
	}(myChan2, wg)
	wg.Wait()
}
