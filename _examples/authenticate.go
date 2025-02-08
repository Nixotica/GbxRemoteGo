package main

import (
	"fmt"

	"github.com/Nixotica/GbxRemoteGo/client"
)

func main() {
	c, err := client.NewClient("127.0.0.1", 5001)
	if err != nil {
		panic(err)
	}

	// Call Authenticate
	res, err := c.Authenticate("SuperAdmin", "SuperAdmin")
	if err != nil {
		panic(err)
	}

	if !res.Success {
		panic("Authentication failed.")
	}

	fmt.Println("Successfully authenticated as SuperAdmin.")
}
