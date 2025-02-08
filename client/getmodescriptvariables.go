package client

// TODO - add back once more knowledge of this
// import (
// 	"encoding/xml"
// 	"fmt"

// 	"github.com/Nixotica/GbxRemoteGo/internal/request"
// 	"github.com/Nixotica/GbxRemoteGo/internal/transport"
// )

// // GetModeScriptVariablesResponse represents the structured response from GetModeScriptVariables
// type GetModeScriptVariablesResponse struct {
// 	Variables []string
// }

// // GetModeScriptVariablesResponse custom XML parsing logic
// func (r *GetModeScriptVariablesResponse) ParseXML(data []byte) error {
// 	fmt.Println("Parsing xml:", string(data))
// 	return xml.Unmarshal(data, &r)
// }

// // GetModeScriptVariables returns the XML-RPC variables available in the current mode script
// func (c *XMLRPCClient) GetModeScriptVariables() (*GetModeScriptVariablesResponse, error) {
// 	return transport.SendXMLRPCRequest(c.Conn, *request.NewGenericRequest("GetModeScriptVariables"), &GetModeScriptVariablesResponse{})
// }
