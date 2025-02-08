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

	res, err := c.Authenticate("SuperAdmin", "SuperAdmin")
	if err != nil {
		panic(err)
	}

	if !res.Success {
		panic("Authentication failed.")
	}

	// Call SetModeScript for CupMode
	_, err = c.SetModeScriptName("Trackmania/TM_TimeAttack_Online")
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated script to TimeAttack.")
}
