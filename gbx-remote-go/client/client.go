package client

import (
	"encoding/binary"
	"fmt"
	"gbx-remote-go/models"
	"gbx-remote-go/transport"
	"log"
	"net"
)

// XMLRPCClient represents a GBXRemote XML-RPC client for Trackmania.
// Docs: https://wiki.trackmania.io/en/dedicated-server/XML-RPC/gbxremote-protocol
type XMLRPCClient struct {
	Host    string
	Port    int
	Conn    net.Conn
	Handler uint32
}

// Result is a generic type for handling async responses.
type Result[T any] struct {
	Value T
	Err   error
}

// NewClient intializes and connects to the server.
func NewClient(host string, port int) (*XMLRPCClient, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}

	client := &XMLRPCClient{
		Host: host,
		Port: port,
		Conn: conn,
	}

	if err = client.validateHeader(); err != nil {
		conn.Close()
		return nil, err
	}

	return client, nil
}

// validateHeader reads and validates the GBXRemote header.
func (c *XMLRPCClient) validateHeader() error {
	headerLenBytes := make([]byte, 4)
	if _, err := c.Conn.Read(headerLenBytes); err != nil {
		return fmt.Errorf("failed to read header length: %v", err)
	}

	headerLen := binary.LittleEndian.Uint32(headerLenBytes)
	headerData := make([]byte, headerLen)
	if _, err := c.Conn.Read(headerData); err != nil {
		return fmt.Errorf("failed to read header data: %v", err)
	}

	if string(headerData) != "GBXRemote 2" {
		return fmt.Errorf("invalid header: %s", headerData)
	}

	log.Println("Connected and validated GBXRemote header.")

	return nil
}

// ListMethods calls system.listMethods and returns available XML-RPC methods. 
func (c *XMLRPCClient) ListMethods() (models.ListMethodsResponse, error) {
	request := models.NewListMethodsRequest()
	response := &models.ListMethodsResponse{}
	response, err := transport.SendXMLRPCRequest(c.Conn, *request, response)
	if err != nil {
		return models.ListMethodsResponse{}, fmt.Errorf("failed to list methods: %v", err)
	}
	return *response, nil
}

// GetStatus calls GetStatus and returns the server status.
func (c *XMLRPCClient) GetStatus() (models.GetStatusResponse, error) {
	request := models.NewGetStatusRequest()
	response := &models.GetStatusResponse{}
	response, err := transport.SendXMLRPCRequest(c.Conn, *request, response)
	if err != nil {
		return models.GetStatusResponse{}, fmt.Errorf("failed to get status: %v", err)
	}
	return *response, nil
}

// GetStatusAsync calls GetStatus and returns a channel to the goroutine return the server status.
func (c *XMLRPCClient) GetStatusAsync() chan Result[*models.GetStatusResponse] {
	resultChan := make(chan Result[*models.GetStatusResponse], 1)

	go func() {
		status, err := c.GetStatus()
		resultChan <- Result[*models.GetStatusResponse]{Value: &status, Err: err}
		close(resultChan)
	}()

	return resultChan
}