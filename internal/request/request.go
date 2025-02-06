package request

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

// GenericRequest is a reusable request model.
type GenericRequest struct {
	Method string        `xmlrpc:"method"`
	Params []interface{} `xmlrpc:"params"`
}

// NewGenericRequest creates a generic XML-RPC request.
func NewGenericRequest(method string, params ...interface{}) *GenericRequest {
	return &GenericRequest{
		Method: method,
		Params: params,
	}
}

// BuildPacket constructs the binary XML-RPC packet for transmission.
func (r *GenericRequest) BuildPacket(handler uint32) ([]byte, error) {
	var xmlPayload bytes.Buffer

	// Start XML-RPC request
	xmlPayload.WriteString(fmt.Sprintf(`<?xml version="1.0"?><methodCall><methodName>%s</methodName><params>`, r.Method))

	// Handle each parameter with the correct XML-RPC format
	for _, param := range r.Params {
		xmlPayload.WriteString("<param><value>")

		switch v := param.(type) {
		case int:
			xmlPayload.WriteString(fmt.Sprintf("<int>%d</int>", v))
		case bool:
			boolVal := "0"
			if v {
				boolVal = "1"
			}
			xmlPayload.WriteString(fmt.Sprintf("<boolean>%s</boolean>", boolVal))
		case string:
			xmlPayload.WriteString(fmt.Sprintf("<string>%s</string>", v))
		case float64:
			xmlPayload.WriteString(fmt.Sprintf("<double>%f</double>", v))
		default:
			return nil, fmt.Errorf("unsupported parameter type: %s", reflect.TypeOf(param))
		}

		xmlPayload.WriteString("</value></param>")
	}

	// Close XML payload
	xmlPayload.WriteString("</params></methodCall>")

	packet := new(bytes.Buffer)

	// Packet length
	packetLen := uint32(xmlPayload.Len())
	binary.Write(packet, binary.LittleEndian, packetLen)

	// Handler
	binary.Write(packet, binary.LittleEndian, handler)

	// XML data
	packet.Write(xmlPayload.Bytes())

	return packet.Bytes(), nil
}
