package ncaa

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
