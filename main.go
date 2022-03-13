package main

import (
	"fmt"
	"scan-port/port"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go port.InitialScan("localhost", &wg)
	go port.WideScan("localhost", &wg)
	wg.Wait()

	fmt.Println("End")
}
