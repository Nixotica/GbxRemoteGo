package main

import (
	"fmt"
	"sync"

	"github.com/Nixotica/GbxRemoteGo/client"
)

func main() {
	c, err := client.NewClient("127.0.0.1", 5001)
	if err != nil {
		panic(err)
	}

	// Call GetStatus
	res, err := c.GetStatus()
	if err != nil {
		panic(err)
	}

	fmt.Println("Server status:", res.Code, res.Name)

	// Call GetStatus async
	var wg sync.WaitGroup
	wg.Add(1)

	getStatusChan := c.GetStatusAsync()
	go func() {
		defer wg.Done()

		result := <-getStatusChan
		if result.Err != nil {
			panic(err)
			return
		}
		fmt.Println("Server Status (async):", result.Value.Code, result.Value.Name)
	}()

	wg.Wait()
}
