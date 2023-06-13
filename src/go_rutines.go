package src

import (
	"fmt"
	"sync"
)

func sayHelloTo(text string, wg *sync.WaitGroup) {
	wg.Done()

	fmt.Println("Hello, ", text)
}

func TestGoRutine() {
	var wg sync.WaitGroup

	wg.Add(1)

	// sayHelloTo("Luis")
	fmt.Println("Hello")

	go sayHelloTo("Daniel", &wg)

	wg.Wait()

	// Not an efficient way to handle routines
	// time.Sleep(time.Second * 1)
}
