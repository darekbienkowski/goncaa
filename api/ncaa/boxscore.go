package ncaa

import (
	"encoding/json"
	"strings"
)

type TeamGameinfoNCAA struct {
}

type GameinfoWrapper struct {
	Contests []GameinfoNCAA `json:"contests"`
}

type GameinfoNCAA struct {
	SportCode         string `json:"sportCode"`
	SportUrl          string `json:"sportUrl"`
	Clock             string `json:"clock"`
	CurrentPeriod     string `json:"currentPeriod"`
	FinalMessage      string `json:"finalMessage"`
	GameState         string `json:"gameState"`
	StatusCodeDisplay string `json:"statusCodeDisplay"`
	StartTime         string `json:"startTime"`
	StartTimeEpoch    int64  `json:"startTimeEpoch"`
	SeasonYear        int    `json:"seasonYear"`
	HasStartTime      bool   `json:"hasStartTime"`
	HasTeamStats      bool   `json:"hasTeamStats"`
	Teams             []GameinfoTeamNCAA
	Location          Location
}

type Location struct {
	Venue     string `json:"venue"`
	City      string `json:"city"`
	StateUsps string `json:"stateUsps"`
}

type GameinfoTeams struct {
	Away GameinfoTeamNCAA
	Home GameinfoTeamNCAA
}

type GameinfoTeamNCAA struct {
	TeamId       string `json:"teamId"`
	IsHome       bool   `json:"isHome"`
	Color        string `json:"color"`
	Seoname      string `json:"seoname"`
	NameFull     string `json:"nameFull"`
	NameShort    string `json:"nameShort"`
	Name6Char    string `json:"name6Char"`
	Score        int    `json:"score"`
	Record       string `json:"record"`
	DivisionName string `json:"divisionName"`
	IsWinner     bool   `json:"isWinner"`
}

type GameinfoPlayerNCAA struct {
	Number               int    `json:"number"`
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	FullName             string `json:"fullName"`
	Position             string `json:"position"`
	MinutesPlayed        int    `json:"minutesPlayed"`
	Starter              bool   `json:"starter"`
	FieldGoalsMade       int    `json:"fieldGoalsMade"`
	FieldGoalsAttempted  int    `json:"fieldGoalsAttempted"`
	FreeThrowsMade       int    `json:"freeThrowsMade"`
	FreeThrowsAttempted  int    `json:"freeThrowsAttempted"`
	ThreePointsMade      int    `json:"threePointsMade"`
	ThreePointsAttempted int    `json:"threePointsAttempted"`
	OffensiveRebounds    int    `json:"offensiveRebounds"`
	TotalRebounds        int    `json:"totalRebounds"`
	Assists              int    `json:"assists"`
	Turnovers            int    `json:"turnovers"`
	PersonalFouls        int    `json:"personalFouls"`
	Steals               int    `json:"steals"`
	BlockedShots         int    `json:"blockedShots"`
	Points               int    `json:"points"`
}

func (p *GameinfoPlayerNCAA) UnmarshalJSON(data []byte) error {
	type alias GameinfoPlayerNCAA
	var aux alias
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	*p = GameinfoPlayerNCAA(aux)
	p.FullName = strings.TrimSpace(aux.FirstName + " " + aux.LastName)
	return nil
}
