package ncaa

import (
	"bytes"
	"strconv"
	"strings"
)

type GameTeamNCAA struct {
	Score IntOrEmpty `json:"score"`
	Names struct {
		Char6 string `json:"char6"`
		Short string `json:"short"`
	} `json:"names"`
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
