package client

import (
	"encoding/xml"
	"fmt"

	"github.com/Nixotica/GbxRemoteGo/internal/request"
	"github.com/Nixotica/GbxRemoteGo/internal/transport"
	"github.com/Nixotica/GbxRemoteGo/models"
)

// GetPlayerListResponse represents the structured response from GetPlayerList.
type GetPlayerListResponse struct {
	Players []models.PlayerInfo
}

// GetPlayerListResponse custom XML parsing logic
func (r *GetPlayerListResponse) ParseXML(data []byte) error {
	// Temporary struct to capture all player <struct> elements
	var temp struct {
		Players []struct {
			Members []struct {
				Name  string `xml:"name"`
				Value struct {
					Int    int     `xml:"i4"`
					Float  float64 `xml:"double"`
					String string  `xml:"string"`
				} `xml:"value"`
			} `xml:"struct>member"`
		} `xml:"params>param>value>array>data>value"`
	}

	// Unmarshal into the temporary struct
	if err := xml.Unmarshal(data, &temp); err != nil {
		return fmt.Errorf("failed to parse XML: %v", err)
	}

	// Convert extracted values into []PlayerInfo
	for _, playerStruct := range temp.Players {
		player := models.PlayerInfo{}
		for _, member := range playerStruct.Members {
			switch member.Name {
			case "Login":
				player.Login = member.Value.String
			case "NickName":
				player.NickName = member.Value.String
			case "PlayerId":
				player.PlayerId = member.Value.Int
			case "TeamId":
				player.TeamId = member.Value.Int
			case "SpectatorStatus":
				player.SpectatorStatus = member.Value.Int
			case "LadderRanking":
				player.LadderRanking = member.Value.Int
			case "Flags":
				player.Flags = member.Value.Int
			case "LadderScore":
				player.LadderScore = member.Value.Float
			}
		}
		r.Players = append(r.Players, player)
	}

	return nil
}

// GetPlayerList retrieves a list of players currently on the server.
//
// Parameters:
//   - maxInfos (int): The maximum number of player entries to return.
//   - startIndex (int): The starting index in the list of players.
//   - structVersion (int): (Optional) Compatibility mode for different game versions:
//       - 0 = Trackmania United
//       - 1 = Trackmania Forever
//       - 2 = Trackmania Forever (including servers)
//
// Returns:
//   - []models.PlayerInfo: A slice of PlayerInfo structures representing players on the server.
//   - error: An error if the request fails.
//
// Notes:
//   - LadderRanking is 0 when not in official mode.
//   - Flags encoding:
//       - ForceSpectator (0,1,2) + StereoDisplayMode * 1000 + IsManagedByAnotherServer * 10000
//       - IsServer * 100000 + HasPlayerSlot * 1000000 + IsBroadcasting * 10000000
//       - HasJoinedGame * 100000000
//   - SpectatorStatus encoding:
//       - Spectator + TemporarySpectator * 10 + PureSpectator * 100
//       - AutoTarget * 1000 + CurrentTargetId * 10000
func (c *XMLRPCClient) GetPlayerList(maxInfos, startIndex, structVersion int) (*GetPlayerListResponse, error) {
	req := request.NewGenericRequest("GetPlayerList", maxInfos, startIndex, structVersion) // TODO - enum struct version
	return transport.SendXMLRPCRequest(c.Conn, *req, &GetPlayerListResponse{})
}
