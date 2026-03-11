package ncaa

import (
	"fmt"
	"time"
)

type GameNCAA struct {
	GameID         int          `json:"gameID,string"`
	StartTime      string       `json:"startTime"`
	StartDate      string       `json:"startDate"`
	StartTimeEpoch int64        `json:"startTimeEpoch,string"`
	GameTitle      string       `json:"title"`
	Away           GameTeamNCAA `json:"away"`
	Home           GameTeamNCAA `json:"home"`
	GameState      string       `json:"gameState"`
}

func (g GameNCAA) FilterValue() string {
	return fmt.Sprintf("%s vs %s", g.Away.Names.Char6, g.Away.Names.Char6)
}

func (g GameNCAA) Title() string {

	t := time.Unix(g.StartTimeEpoch, 0)
	dateString := t.Format("Jan 02, 3:04 PM")

	return fmt.Sprintf("%s (%d) vs %s (%d) - (%s)",
		g.Away.Names.Short,
		g.Away.Score,
		g.Home.Names.Short,
		g.Home.Score,
		dateString,
	)
}

func (g GameNCAA) Description() string {
	// TODO figure something out to put here
	return g.GameState
}
