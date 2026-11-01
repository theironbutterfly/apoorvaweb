package main

import (
	"fmt"
	"sync"
)

func reboot(serverID int, wg *sync.WaitGroup) {

	// TODO 1:
	// Signal when this goroutine finishes.
	wg.Done()

	fmt.Println("Rebooting server", serverID)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {

		wg.Add(1)

		go reboot(i, &wg)
	}

	// TODO 3:
	// Wait until all tasks complete.
	wg.Wait()

	fmt.Println("All servers rebooted")
}

// Three tools you can use
// Go's sync.WaitGroup provides three important methods.
// Method	Purpose
// wg.Add(1)	Register one unit of work
// wg.Done()	Signal that one unit of work finished
// wg.Wait()	Block until all registered work finishes
