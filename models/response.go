package models

import (
	"encoding/xml"
)

// XMLRPCResponse is a generic XML-RPC response wrapper.
type XMLRPCResponse struct {
	XMLName xml.Name	`xml:"methodResponse"`
}

// ParseXMLResponse unmarshals the XML response into the provided response struct.
func ParseXMLResponse[T any](responseData []byte, responseStruct *T) (*T, error) {
	err := xml.Unmarshal(responseData, responseStruct)
	if err != nil {
		return nil, err
	}

	return responseStruct, nil
}

// ListMethodsResponse represents the structured response from system.listMethods
type ListMethodsResponse struct {
	Methods []string `xml:"params>param>value>array>data>value>string"`
}

// GetStatusResponse represents the structured response from GetStatus
type GetStatusResponse struct {
	Status struct {
		Code int	`xml:"member>value>i4"` // TODO map codes to enum
		Name string	`xml:"member>value>string"`
	} `xml:"params>param>value>struct"`
}