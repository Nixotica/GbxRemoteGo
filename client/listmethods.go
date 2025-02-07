package client

import (
	"encoding/xml"

	"github.com/Nixotica/GbxRemoteGo/internal/request"
	"github.com/Nixotica/GbxRemoteGo/internal/transport"
)

// ListMethodsResponse represents the structured response from system.listMethods
type ListMethodsResponse struct {
	Methods []string `xml:"params>param>value>array>data>value>string"`
}

// LisetMethodsResponse custom XML parsing logic
func (r *ListMethodsResponse) ParseXML(data []byte) error {
	return xml.Unmarshal(data, &r)
}

// ListMethods calls system.listMethods and returns available XML-RPC methods.
func (c *XMLRPCClient) ListMethods() (*ListMethodsResponse, error) {
	return transport.SendXMLRPCRequest(c.Conn, *request.NewGenericRequest("system.listMethods"), &ListMethodsResponse{})
}
