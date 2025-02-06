package main

import (
	"fmt"
	"gbx-remote-go/client"
	"log"
)

func main() {
	xmlRpcClient, err := client.NewClient("127.0.0.1", 5001)
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	defer xmlRpcClient.Conn.Close()

	// Call system.listMethods 
	listMethodsResponse, err := xmlRpcClient.ListMethods()
	if err != nil {
		log.Fatalf("Error calling method: %v", err)
	}

	fmt.Println("Available XML-RPC methods:")
	for _, method := range listMethodsResponse.Methods {
		fmt.Println("-", method)
	}

	// Call GetStatus
	getStatusResponse, err := xmlRpcClient.GetStatus()
	if err != nil {
		log.Fatalf("Error calling method: %v", err)
	}
	fmt.Println("Server Status:", getStatusResponse.Status.Code, getStatusResponse.Status.Name)
}