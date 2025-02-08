package transport

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/Nixotica/GbxRemoteGo/internal/request"
	"github.com/Nixotica/GbxRemoteGo/internal/response"
)

// SendXMLRPCRequest sends an XML-RPC request and unmarshals the response into the given response struct.
func SendXMLRPCRequest[T response.GenericResponse](conn net.Conn, request request.GenericRequest, responseStruct T) (T, error) {
	// Construct packet
	handler := getNextHandler()
	packetBytes, err := request.BuildPacket(handler)
	if err != nil {
		return responseStruct, fmt.Errorf("failed to build packet: %v", err)
	}

	// Send request
	_, err = conn.Write(packetBytes)
	if err != nil {
		return responseStruct, fmt.Errorf("failed to send request: %v", err)
	}

	// Read response header (8 bytes)
	responseHeader := make([]byte, 8)
	_, err = conn.Read(responseHeader)
	if err != nil {
		return responseStruct, fmt.Errorf("failed to read response header: %v", err)
	}

	// Get response size
	responseSize := binary.LittleEndian.Uint32(responseHeader[:4])

	// Get response data
	responseData := make([]byte, responseSize)
	_, err = conn.Read(responseData)
	if err != nil {
		return responseStruct, fmt.Errorf("failed to read response data: %v", err)
	}

	// Parse response XML into provided response struct
	err = responseStruct.ParseXML(responseData)
	if err != nil {
		return responseStruct, fmt.Errorf("failed to parse response XML: %v", err)
	}

	return responseStruct, nil
}
