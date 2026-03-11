package repositories

import (
	"encoding/json"
	"time"

	"github.com/darekbienkowski/goncaa/api/ncaa"
)

var scheduleRepoInstance *ScheduleRepository

type ScheduleRepository struct {
	client *ncaa.Client
}

func NewScheduleRepository() *ScheduleRepository {
	if scheduleRepoInstance == nil {
		scheduleRepoInstance = &ScheduleRepository{
			client: ncaa.NewDefaultClient(),
		}
	}

	return scheduleRepoInstance
}

func (repo *ScheduleRepository) GetScheduleForDate(date time.Time) (*ncaa.Schedule, error) {
	queryParams := make(map[string]string)

	urlDate := date.Format("2006/01/02/")
	urlPath := "scoreboard/basketball-men/d1/" + urlDate + "all-conf"

	responseBytes, err := repo.client.Get(urlPath, queryParams)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response
	var res ncaa.Schedule
	if err := json.Unmarshal(responseBytes, &res); err != nil {
		panic(err)
	}

	return &res, nil

}
