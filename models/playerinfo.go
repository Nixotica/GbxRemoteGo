package models

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
