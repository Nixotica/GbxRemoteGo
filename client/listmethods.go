package client

import (
	"fmt"

	"github.com/Nixotica/GbxRemoteGo/internal/transport"
	"github.com/Nixotica/GbxRemoteGo/internal/request"
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
	if err != nil {
		return ListMethodsResponse{}, fmt.Errorf("failed to list methods: %v", err)
	}
	return *res, nil
}
