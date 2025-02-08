package client

import (
	"encoding/xml"

	"github.com/Nixotica/GbxRemoteGo/internal/request"
	"github.com/Nixotica/GbxRemoteGo/internal/transport"
)

/* SetModeScriptText */

// SetModeScriptResponse represents the structured response from SetModeScriptText
type SetModeScriptResponse struct{}

// SetModeScriptResponse custom XML parsing logic
func (r *SetModeScriptResponse) ParseXML(data []byte) error {
	return xml.Unmarshal(data, &r)
}

// SetModeScriptText calls SetModeScriptText with the entire script as a string and returns True on success.
func (c *XMLRPCClient) SetModeScriptText(scriptText string) (bool, error) {
	_, _ = transport.SendXMLRPCRequest(c.Conn, *request.NewGenericRequest("SetModeScriptText", scriptText), &SetModeScriptResponse{})
	return true, nil
}

/* GetModeScriptText */

// SetModeScriptName calls SetModeScriptName with the name of a script as a string and returns True on success.
func (c *XMLRPCClient) SetModeScriptName(scriptName string) (bool, error) {
	_, _ = transport.SendXMLRPCRequest(c.Conn, *request.NewGenericRequest("SetModeScriptName", scriptName), &SetModeScriptResponse{})
	return true, nil
}
