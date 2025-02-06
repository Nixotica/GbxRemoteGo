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

	// Call system.ListMethods
	res, err := c.ListMethods()
	if err != nil {
		panic(err)
	}

	fmt.Println("Methods:")
	for _, method := range res.Methods {
		fmt.Println("-", method)
	}
}
