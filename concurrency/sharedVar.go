package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		x := i
		go func() {
			defer wg.Done()
			fmt.Println(&x)
		}()
	}
	wg.Wait()
	fmt.Println("Done")
}
