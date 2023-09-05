// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func parseFile(parsedLogsCh chan int, filesWg *sync.WaitGroup, itemsWg *sync.WaitGroup, filePath string) {
	defer filesWg.Done()
	fmt.Println(filePath)

	for i := 0; i < 100; i++ {
		itemsWg.Add(1)
		// parse line
		parsedLogsCh <- i
	}
}

func main() {
	filePaths := []string{"XXX", "YYY", "ZZZ"}

	filesWg := new(sync.WaitGroup)
	filesWg.Add(len(filePaths))

	parsedLogsCh := make(chan int, 2)
	itemsWg := new(sync.WaitGroup)

	// 1. parse file
	for _, filePath := range filePaths {
		go parseFile(parsedLogsCh, filesWg, itemsWg, filePath)
	}

	go func(parsedLogsCh chan int, itemsWg *sync.WaitGroup) {
		for {
			val := <-parsedLogsCh
			// add to list
			fmt.Println(val)
			itemsWg.Done()
		}
	}(parsedLogsCh, itemsWg)

	filesWg.Wait()
	itemsWg.Wait()

	// 2. render template

	fmt.Println("Hello, 世界")
}
