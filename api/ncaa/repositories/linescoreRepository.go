package repositories

import (
	"encoding/json"
	"fmt"

	"github.com/darekbienkowski/goncaa/api/ncaa"
)

type LinescoreRepository struct {
	ncaaClient *ncaa.Client
}

func NewLinescoreRepository() LinescoreRepository {
	return LinescoreRepository{
		ncaaClient: ncaa.NewDefaultClient(),
	}
}

func (repo LinescoreRepository) GetLinescoreFromGameId(gameId int) (*ncaa.Linescore, error) {
	responseBytes, err := repo.ncaaClient.Get(fmt.Sprintf("game/%d/boxscore", gameId), nil)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response
	var res ncaa.Linescore
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
