package ncaa

import (
	"encoding/json"
	"strings"
)

type Linescore struct {
	Title         string          `json:"title"`
	ContestId     int             `json:"contestId"`
	Description   string          `json:"description"`
	DivisionName  string          `json:"divisionName"`
	Status        string          `json:"status"`
	Period        string          `json:"period"`
	Minutes       IntOrEmpty      `json:"minutes"`
	Seconds       IntOrEmpty      `json:"seconds"`
	SportCode     string          `json:"sportCode"`
	Teams         []LinescoreTeam `json:"teams"`
	TeamBoxscores []TeamBoxscore  `json:"teamBoxscore"`
}

type LinescoreTeam struct {
	IsHome    bool   `json:"isHome"`
	TeamID    string `json:"teamId"`
	SeoName   string `json:"seoname"`
	Name6     string `json:"name6Char"`
	NameFull  string `json:"nameFull"`
	NameShort string `json:"nameShort"`
	TeamName  string `json:"teamName"`
	Color     string `json:"color"`
}

type TeamBoxscore struct {
	TeamID      int                     `json:"teamId"`
	PlayerStats []PlayerStatsBasketball `json:"playerStats"`
	TeamStats   TeamStatsBasketball     `json:"teamStats"`
}

type PlayerStatsBasketball struct {
	ID                   int    `json:"id"`
	Number               int    `json:"number"`
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	FullName             string `json:"fullName"`
	Position             string `json:"position"`
	MinutesPlayed        string `json:"minutesPlayed"`
	Starter              bool   `json:"starter"`
	FieldGoalsMade       string `json:"fieldGoalsMade"`
	FieldGoalsAttempted  string `json:"fieldGoalsAttempted"`
	FreeThrowsMade       string `json:"freeThrowsMade"`
	FreeThrowsAttempted  string `json:"freeThrowsAttempted"`
	ThreePointsMade      string `json:"threePointsMade"`
	ThreePointsAttempted string `json:"threePointsAttempted"`
	OffensiveRebounds    string `json:"offensiveRebounds"`
	TotalRebounds        string `json:"totalRebounds"`
	Assists              string `json:"assists"`
	Turnovers            string `json:"turnovers"`
	PersonalFouls        string `json:"personalFouls"`
	Steals               string `json:"steals"`
	BlockedShots         string `json:"blockedShots"`
	Points               string `json:"points"`
}

func (p *PlayerStatsBasketball) UnmarshalJSON(data []byte) error {
	type alias PlayerStatsBasketball
	var aux alias
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	*p = PlayerStatsBasketball(aux)
	p.FullName = strings.TrimSpace(aux.FirstName + " " + aux.LastName)
	return nil
}

type TeamStatsBasketball struct {
	FieldGoalsMade       string `json:"fieldGoalsMade"`
	FieldGoalsAttempted  string `json:"fieldGoalsAttempted"`
	FreeThrowsMade       string `json:"freeThrowsMade"`
	FreeThrowsAttempted  string `json:"freeThrowsAttempted"`
	ThreePointsMade      string `json:"threePointsMade"`
	ThreePointsAttempted string `json:"threePointsAttempted"`
	OffensiveRebounds    string `json:"offensiveRebounds"`
	TotalRebounds        string `json:"totalRebounds"`
	Assists              string `json:"assists"`
	Turnovers            string `json:"turnovers"`
	PersonalFouls        string `json:"personalFouls"`
	Steals               string `json:"steals"`
	BlockedShots         string `json:"blockedShots"`
	Points               string `json:"points"`
	FieldGoalPercentage  string `json:"fieldGoalPercentage"`
	ThreePointPercentage string `json:"threePointPercentage"`
	FreeThrowPercentage  string `json:"freeThrowPercentage"`
}
