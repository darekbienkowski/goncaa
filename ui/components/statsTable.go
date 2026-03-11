package components

import (
	"fmt"

	"github.com/darekbienkowski/goncaa/api/ncaa"

	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

func playerTableColumns(maxPlayerLength int) []table.Column {
	return []table.Column{
		table.NewColumn("playerName", "Player", maxPlayerLength+1),
		table.NewColumn("pos", "POS", 3).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("min", "MIN", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("fgm-a", "FGM-A", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("3pm-a", "3PM-A", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("ftm-a", "FTM-A", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("oreb", "OREB", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("reb", "REB", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("ast", "AST", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("st", "ST", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("blk", "BLK", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("to", "TO", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("pf", "PF", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
		table.NewFlexColumn("pts", "PTS", 1).WithStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)),
	}
}

func BuildPlayerStatsTable(players []ncaa.PlayerStatsBasketball) table.Model {
	playerNameMaxLen := getPlayerNameMaxLen(players)
	tableColumns := playerTableColumns(playerNameMaxLen)

	lineup := []table.Row{}

	for _, player := range players {
		lineup = append(lineup, boxscorePlayerToPlayerTableRow(player))
	}

	return table.New(tableColumns).WithRows(lineup).WithBaseStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Left))
}

func getPlayerNameMaxLen(players []ncaa.PlayerStatsBasketball) int {
	playerNameMaxLen := 0
	for _, player := range players {
		if len(player.FullName) > playerNameMaxLen {
			playerNameMaxLen = len(player.FullName)
		}
	}

	return playerNameMaxLen
}

func boxscorePlayerToPlayerTableRow(player ncaa.PlayerStatsBasketball) table.Row {
	row := table.NewRow(table.RowData{
		"id":         player.ID,
		"playerName": player.FullName,
		"pos":        player.Position,
		"min":        player.MinutesPlayed,
		"fgm-a":      fmt.Sprintf("%s-%s", player.FieldGoalsMade, player.FieldGoalsAttempted),
		"3pm-a":      fmt.Sprintf("%s-%s", player.ThreePointsMade, player.ThreePointsAttempted),
		"ftm-a":      fmt.Sprintf("%s-%s", player.FreeThrowsMade, player.FreeThrowsAttempted),
		"oreb":       player.OffensiveRebounds,
		"reb":        player.TotalRebounds,
		"ast":        player.Assists,
		"st":         player.Steals,
		"blk":        player.BlockedShots,
		"to":         player.Turnovers,
		"pf":         player.PersonalFouls,
		"pts":        player.Points,
	})

	return row
}

func EmptyPlayerStatsTable() table.Model {
	return table.New(playerTableColumns(10)).WithRows([]table.Row{}).WithBaseStyle(lipgloss.NewStyle().AlignHorizontal(lipgloss.Left))
}
