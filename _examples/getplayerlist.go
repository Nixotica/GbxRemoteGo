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

	// Call GetPlayerList (for 4 players, 0 offset, forever-server mode)
	res, err := c.GetPlayerList(4, 0, 2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Players:")
	for _, player := range res.Players {
		fmt.Println("-", player)
	}
}
