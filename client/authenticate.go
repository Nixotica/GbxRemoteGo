package client

import (
	"encoding/xml"

	"github.com/Nixotica/GbxRemoteGo/internal/request"
	"github.com/Nixotica/GbxRemoteGo/internal/response"
	"github.com/Nixotica/GbxRemoteGo/internal/transport"
)

type AuthenticateResponse struct {
	Success bool `xml:"params>param>value>boolean"`
}

func (r *AuthenticateResponse) ParseXML(data []byte) error {
	var faultWrapper struct {
		Fault struct {
			Members []struct {
				Name  string `xml:"name"`
				Value struct {
					Int    int    `xml:"int"`
					String string `xml:"string"`
				} `xml:"value"`
			} `xml:"member"`
		} `xml:"fault>value>struct"`
	}

	if err := xml.Unmarshal(data, &faultWrapper); err == nil && len(faultWrapper.Fault.Members) > 0 {
		// Extract fault details
		var fault response.FaultResponse
		for _, member := range faultWrapper.Fault.Members {
			switch member.Name {
			case "faultCode":
				fault.Code = member.Value.Int
			case "faultString":
				fault.Message = member.Value.String
			}
		}
		return fault // Return FaultResponse as an error
	}

	// Otherwise return normal response
	return xml.Unmarshal(data, &r)
}

// Authenticate allows caller to specify a login and password to gain access to a set of
// functionalities corresponding to this authorization level.
func (c *XMLRPCClient) Authenticate(login string, password string) (*AuthenticateResponse, error) {
	req := request.NewGenericRequest("Authenticate", login, password)
	return transport.SendXMLRPCRequest(c.Conn, *req, &AuthenticateResponse{})
}
