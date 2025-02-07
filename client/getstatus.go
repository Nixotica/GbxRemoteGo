package client

import (
	"encoding/xml"
	"fmt"

	"github.com/Nixotica/GbxRemoteGo/internal/request"
	"github.com/Nixotica/GbxRemoteGo/internal/transport"
)

// GetStatusResponse represents the structured response from GetStatus
type GetStatusResponse struct {
	Code int
	Name string
}

// GetStatusResponse custom XML parsing logic
func (r *GetStatusResponse) ParseXML(data []byte) error {
	// Temporary struct to capture <member> elements dynamically
	var temp struct {
		Members []struct {
			Name  string `xml:"name"`
			Value struct {
				Int    int    `xml:"i4"`
				String string `xml:"string"`
			} `xml:"value"`
		} `xml:"params>param>value>struct>member"`
	}

	// Unmarshal the XML into the temporary struct
	if err := xml.Unmarshal(data, &temp); err != nil {
		return fmt.Errorf("failed to parse XML: %v", err)
	}

	// Map extracted values into GetStatusResponse fields
	for _, member := range temp.Members {
		switch member.Name {
		case "Code":
			r.Code = member.Value.Int
		case "Name":
			r.Name = member.Value.String
		}
	}

	return nil
}

// GetStatus calls GetStatus and returns the server status.
func (c *XMLRPCClient) GetStatus() (*GetStatusResponse, error) {
	return transport.SendXMLRPCRequest(c.Conn, *request.NewGenericRequest("GetStatus"), &GetStatusResponse{})
}

// GetStatusAsync calls GetStatus and returns a channel to the goroutine return the server status.
func (c *XMLRPCClient) GetStatusAsync() chan Result[*GetStatusResponse] {
	resultChan := make(chan Result[*GetStatusResponse], 1)

	go func() {
		status, err := c.GetStatus()
		resultChan <- Result[*GetStatusResponse]{Value: status, Err: err}
		close(resultChan)
	}()

	return resultChan
}
