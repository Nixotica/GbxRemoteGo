package client

import (
	"github.com/Nixotica/GbxRemoteGo/internal/request"
	"github.com/Nixotica/GbxRemoteGo/internal/transport"
)

// ListMethodsResponse represents the structured response from system.listMethods
type ListMethodsResponse struct {
	Methods []string `xml:"params>param>value>array>data>value>string"`
}

// ListMethods calls system.listMethods and returns available XML-RPC methods.
func (c *XMLRPCClient) ListMethods() (ListMethodsResponse, error) {
	req := request.NewGenericRequest("system.listMethods")
	res := &ListMethodsResponse{}
	res, err := transport.SendXMLRPCRequest(c.Conn, *req, res)
	return *res, err
}
