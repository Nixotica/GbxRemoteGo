package client

import (
	"encoding/binary"
	"fmt"
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

