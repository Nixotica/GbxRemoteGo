package response

import "fmt"

// FaultResponse represents an XML-RPC fault response from the server.
type FaultResponse struct {
	Code    int    `xmlrpc:"faultCode"`
	Message string `xmlrpc:"faultString"`
}

// Error implements the error interface for FaultResponse
func (f FaultResponse) Error() string {
	return fmt.Sprintf("RPC Error: %d: %s", f.Code, f.Message)
}
