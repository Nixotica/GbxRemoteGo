package models

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// GenericRequest is a reusable request model.
type GenericRequest struct {
	Method  		string   `xml:"methodName"`
	Params  		[]string `xml:"params>param>value"`
}

// NewGenericRequest creates a generic XML-RPC request.
func NewGenericRequest(method string, params ...string) *GenericRequest {
	return &GenericRequest{
		Method: method,
		Params: params,
	}
}

// BuildPacket constructs the binary packet for transmission.
func (r *GenericRequest) BuildPacket(handler uint32) ([]byte, error) {
	xmlPayload := fmt.Sprintf(`<?xml version="1.0"?><methodCall><methodName>%s</methodName><params>`, r.Method)
	for _, param := range r.Params {
		xmlPayload += fmt.Sprintf(`<param><value>%s</value></param>`, param)
	}
	xmlPayload += `</params></methodCall>`

	packet := new(bytes.Buffer)

	// Packet length
	packetLen := uint32(len(xmlPayload))
	binary.Write(packet, binary.LittleEndian, packetLen)

	// Handler
	binary.Write(packet, binary.LittleEndian, handler)

	// XML data
	packet.Write([]byte(xmlPayload))

	return packet.Bytes(), nil
}

// NewListMethodsRequest creates a request for system.listMethods
func NewListMethodsRequest() *GenericRequest {
	return NewGenericRequest("system.listMethods")
}

// NewGetStatusRequest creates a request for GetStatus
func NewGetStatusRequest() *GenericRequest {
	return NewGenericRequest("GetStatus")
}

