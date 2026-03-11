package repositories

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/darekbienkowski/goncaa/api/ncaa"
)

type GameinfoRepository struct {
	ncaaClient *ncaa.Client
}

func NewGameinfoRepository() GameinfoRepository {
	return GameinfoRepository{
		ncaaClient: ncaa.NewDefaultClient(),
	}
}

func (repo GameinfoRepository) GetGameinfoFromGameID(gameId int) (*ncaa.GameinfoNCAA, error) {
	responseBytes, err := repo.ncaaClient.Get(fmt.Sprintf("game/%d", gameId), nil)
	if err != nil {
		return nil, err
	}

	var res ncaa.GameinfoWrapper
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		return nil, err
	}

	if len(res.Contests) == 0 {
		return nil, fmt.Errorf("no contests returned for game %d", gameId)
	}

	boxscore := res.Contests[0]
	sortTeamsByIsHome(&boxscore)

	return &boxscore, nil
}

func sortTeamsByIsHome(boxscore *ncaa.GameinfoNCAA) {
	sort.Slice(boxscore.Teams, func(i, j int) bool {
		return !boxscore.Teams[i].IsHome && boxscore.Teams[j].IsHome
	})
}
