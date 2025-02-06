package client

import (
	"github.com/Nixotica/GbxRemoteGo/internal/request"
	"github.com/Nixotica/GbxRemoteGo/internal/transport"
	"github.com/Nixotica/GbxRemoteGo/models"
)

// GetPlayerListResponse represents the structured response from GetPlayerList.
type GetPlayerListResponse struct {
	Players []models.PlayerInfo `xml:"params>param>value>array>data>value>struct"`
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
func (c *XMLRPCClient) GetPlayerList(maxInfos, startIndex, structVersion int) (GetPlayerListResponse, error) {
	req := request.NewGenericRequest("GetPlayerList", maxInfos, startIndex, structVersion) // TODO - enum struct version
	res := &GetPlayerListResponse{}
	res, err := transport.SendXMLRPCRequest(c.Conn, *req, res)
	return *res, err
}
