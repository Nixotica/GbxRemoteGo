package models

import (
	"encoding/xml"
	"fmt"
)

// Member represents a key-value pair in XML-RPC response
type Member struct {
	Name  string `xml:"name"`
	Value struct {
		StringValue string  `xml:"string,omitempty"`
		IntValue    int     `xml:"i4,omitempty"`
		FloatValue  float64 `xml:"double,omitempty"`
	} `xml:"value"`
}

// PlayerInfo represents the details of a player on the server.
type PlayerInfo struct {
	Login           string
	NickName        string
	PlayerId        int
	TeamId          int
	SpectatorStatus int
	LadderRanking   int
	Flags           int
	LadderScore     float64
}

// UnmarshalXML performs a custom unmarshal of the XML returned by the server to GetPlayerListResponse format.
func (p *PlayerInfo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var members []Member
	if err := d.DecodeElement(&members, &start); err != nil {
		return err
	}

	// Map XML-RPC members into PlayerInfo fields
	for _, m := range members {
		switch m.Name {
		case "Login":
			p.Login = m.Value.StringValue
		case "NickName":
			p.NickName = m.Value.StringValue
		case "PlayerId":
			p.PlayerId = m.Value.IntValue
		case "TeamId":
			p.TeamId = m.Value.IntValue
		case "SpectatorStatus":
			p.SpectatorStatus = m.Value.IntValue
		case "LadderRanking":
			p.LadderRanking = m.Value.IntValue
		case "Flags":
			p.Flags = m.Value.IntValue
		case "LadderScore":
			p.LadderScore = m.Value.FloatValue
		default:
			return fmt.Errorf("unknown field: %s", m.Name)
		}
	}

	return nil
}
