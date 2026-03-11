package ncaa

import (
	"bytes"
	"strconv"
	"strings"
)

type GameTeam struct {
	Score        int
	LeagueRecord TeamRecord
	Team         struct {
		Name string
	}
}

type GameTeamNCAA struct {
	Score IntOrEmpty `json:"score"`
	Names struct {
		Char6 string `json:"char6"`
		Short string `json:"short"`
	} `json:"names"`
}

type TeamRecord struct {
	Wins   int
	Losses int
	Pct    string
}

type Team struct {
	Id   int
	Name string
}

type IntOrEmpty int

func (i *IntOrEmpty) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		*i = 0
		return nil
	}
	str := strings.Trim(string(data), `"`)
	if str == "" {
		*i = 0
		return nil
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	*i = IntOrEmpty(val)
	return nil
}
