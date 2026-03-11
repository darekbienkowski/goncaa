package components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/darekbienkowski/goncaa/api/ncaa"

	"github.com/evertras/bubble-table/table"
)

func generateInningsColumns(num int) []table.Column {
	inningsColumns := []table.Column{}

	for i := range num {
		inningsColumns = append(inningsColumns, table.NewFlexColumn(fmt.Sprintf("%d", i+1), fmt.Sprintf("%d", i+1), 1))
	}

	return inningsColumns
}

func linescoreColumns() []table.Column {
	return []table.Column{
		table.NewFlexColumn("teamName", "Team", 3),
		table.NewFlexColumn("fgm-a", "FGM-A", 1),
		table.NewFlexColumn("3pm-a", "3PM-A", 1),
		table.NewFlexColumn("ftm-a", "FTM-A", 1),
		table.NewFlexColumn("oreb", "OREB", 1),
		table.NewFlexColumn("reb", "REB", 1),
		table.NewFlexColumn("ast", "AST", 1),
		table.NewFlexColumn("st", "ST", 1),
		table.NewFlexColumn("blk", "BLK", 1),
		table.NewFlexColumn("to", "TO", 1),
		table.NewFlexColumn("pf", "PF", 1),
		table.NewFlexColumn("pts", "PTS", 1),
	}
}

func createLinescoreRow(team ncaa.LinescoreTeam, teamStat ncaa.TeamStatsBasketball) table.Row {

	row := table.NewRow(table.RowData{
		"teamName": fmt.Sprintf("%s %s", team.NameShort, team.TeamName),
		"fgm-a":    fmt.Sprintf("%s-%s", teamStat.FieldGoalsMade, teamStat.FieldGoalsAttempted),
		"3pm-a":    fmt.Sprintf("%s-%s", teamStat.ThreePointsMade, teamStat.ThreePointsAttempted),
		"ftm-a":    fmt.Sprintf("%s-%s", teamStat.FreeThrowsMade, teamStat.FreeThrowsAttempted),
		"oreb":     teamStat.OffensiveRebounds,
		"reb":      teamStat.TotalRebounds,
		"ast":      teamStat.Assists,
		"st":       teamStat.Steals,
		"blk":      teamStat.BlockedShots,
		"to":       teamStat.Turnovers,
		"pf":       teamStat.PersonalFouls,
		"pts":      teamStat.Points,
	})

	return row
}

func BuildLinescoreTable(awayTeamName string, homeTeamName string, linescore *ncaa.Linescore) table.Model {
	linescoreColumns := linescoreColumns()

	linescoreRows := []table.Row{}

	for i := len(linescore.TeamBoxscores) - 1; i >= 0; i-- {
		teamBoxStat := linescore.TeamBoxscores[i]
		linescoreRows = append(linescoreRows, createLinescoreRow(linescore.Teams[i], teamBoxStat.TeamStats))
	}

	return table.New(linescoreColumns).WithRows(linescoreRows).WithBaseStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center))
}

func EmptyLinescoreTable() table.Model {
	return table.New(linescoreColumns()).WithRows([]table.Row{}).WithBaseStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center))
}
