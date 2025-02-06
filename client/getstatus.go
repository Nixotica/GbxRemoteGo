package client

import (
	"github.com/Nixotica/GbxRemoteGo/internal/request"
	"github.com/Nixotica/GbxRemoteGo/internal/transport"
)

// GetStatusResponse represents the structured response from GetStatus
type GetStatusResponse struct {
	Status struct {
		Code int    `xml:"member>value>i4"` // TODO map codes to enum
		Name string `xml:"member>value>string"`
	} `xml:"params>param>value>struct"`
}

// GetStatus calls GetStatus and returns the server status.
func (c *XMLRPCClient) GetStatus() (GetStatusResponse, error) {
	req := request.NewGenericRequest("GetStatus")
	res := &GetStatusResponse{}
	res, err := transport.SendXMLRPCRequest(c.Conn, *req, res)
	return *res, err
}

// GetStatusAsync calls GetStatus and returns a channel to the goroutine return the server status.
func (c *XMLRPCClient) GetStatusAsync() chan Result[*GetStatusResponse] {
	resultChan := make(chan Result[*GetStatusResponse], 1)

	go func() {
		status, err := c.GetStatus()
		resultChan <- Result[*GetStatusResponse]{Value: &status, Err: err}
		close(resultChan)
	}()

	return resultChan
}
