package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go sayHello()
	wg.Wait()

	var wg2 sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		fmt.Printf("-- %s --\n", salutation)
		wg2.Add(1)
		//go func(salutation string) {
		//	defer wg2.Done()
		//	fmt.Println(salutation)
		//}(salutation)
		go func() {
			defer wg2.Done()
			fmt.Println(salutation)
		}()
	}
	fmt.Println("wait!")
	wg2.Wait()
}

func sayHello() {
	defer wg.Done()
	fmt.Println("hello!")
}