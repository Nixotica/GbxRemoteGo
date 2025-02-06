package transport

import (
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"log"
	"net"

	"github.com/Nixotica/GbxRemoteGo/internal/request"
)

// parseXMLResponse unmarshals the XML response into the provided response struct.
func parseXMLResponse[T any](responseData []byte, responseStruct *T) (*T, error) {
	log.Printf("parseXMLResponse: %s", string(responseData)) // TODO - remove for major versions

	err := xml.Unmarshal(responseData, responseStruct)
	if err != nil {
		return nil, err
	}

	log.Printf("unmarshaled response: %+v", responseStruct) // TODO - remove for major versions

	return responseStruct, nil
}

// SendXMLRPCRequest sends an XML-RPC request and unmarshals the response into the given response struct.
func SendXMLRPCRequest[T any](conn net.Conn, request request.GenericRequest, responseStruct *T) (*T, error) {
	log.Println("Making request to server:", request)

	// Construct packet
	handler := getNextHandler()
	packetBytes, err := request.BuildPacket(handler)
	if err != nil {
		return nil, fmt.Errorf("failed to build packet: %v", err)
	}

	// Send request
	_, err = conn.Write(packetBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	// Read response header (8 bytes)
	responseHeader := make([]byte, 8)
	_, err = conn.Read(responseHeader)
	if err != nil {
		return nil, fmt.Errorf("failed to read response header: %v", err)
	}

	// Get response size
	responseSize := binary.LittleEndian.Uint32(responseHeader[:4])

	// Get response data
	responseData := make([]byte, responseSize)
	_, err = conn.Read(responseData)
	if err != nil {
		return nil, fmt.Errorf("failed to read response data: %v", err)
	}

	// Parse response XML into provided response struct
	parsedResponse, err := parseXMLResponse(responseData, responseStruct)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response XML: %v", err)
	}

	return parsedResponse, nil
}
