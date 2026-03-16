package ui

import (
	"fmt"
	"time"

	"github.com/darekbienkowski/goncaa/api/ncaa"
	"github.com/darekbienkowski/goncaa/api/ncaa/repositories"
	"github.com/darekbienkowski/goncaa/ui/components"
	"github.com/darekbienkowski/goncaa/ui/constants"
	"github.com/darekbienkowski/goncaa/ui/popup"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

var TABLE_TO_INDEX_MAP = map[string]int{
	"awayBatters": 0,
	"homeBatters": 1,
}

const RESPONSIVE_WIDTH_BREAKPOINT = 130
const REFRESH_GAME_RATE = 5

type GameScreenModel struct {
	linescoreTable table.Model
	playerTables   []table.Model
	game           ncaa.GameNCAA
	boxscore       ncaa.GameinfoNCAA
	help           help.Model
	viewport       viewport.Model
	previousModel  Model
	popup          popup.IPopup
	width, height  int
	lastUpdate     time.Time // New feat
}

var gameScreenKM = GameScreenKM{
	Back: key.NewBinding(
		key.WithKeys("esc", "q"),
		key.WithHelp("esc/q", "Back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "Up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "Down"),
	),
	PageUp: key.NewBinding(
		key.WithKeys("pgup", "b"),
		key.WithHelp("pgup/b", "PgUp"),
	),
	PageDown: key.NewBinding(
		key.WithKeys("pgdown", " ", "f"),
		key.WithHelp("pgdown/spc", "PgDn"),
	),
}

func (m GameScreenModel) Init() tea.Cmd {
	return tea.Every(REFRESH_GAME_RATE*time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m GameScreenModel) refreshGameData() (*GameScreenModel, error) {

	boxscore, err := repositories.NewGameinfoRepository().GetGameinfoFromGameID(m.game.GameID)
	if err == nil {
		m.boxscore = *boxscore
	}

	linescore, err := repositories.NewLinescoreRepository().GetLinescoreFromGameId(m.game.GameID)
	if err == nil {

		// Responsive Layout Logic
		var tableWidth int
		if m.width < RESPONSIVE_WIDTH_BREAKPOINT {
			tableWidth = m.width
		} else {
			tableWidth = m.width / 2
		}

		m.linescoreTable = components.BuildLinescoreTable(m.game.Away.Names.Short, m.game.Home.Names.Short, linescore).WithTargetWidth(m.width)
		m.playerTables[TABLE_TO_INDEX_MAP["awayBatters"]] = components.BuildPlayerStatsTable(linescore.TeamBoxscores[1].PlayerStats).WithTargetWidth(tableWidth)
		m.playerTables[TABLE_TO_INDEX_MAP["homeBatters"]] = components.BuildPlayerStatsTable(linescore.TeamBoxscores[0].PlayerStats).WithTargetWidth(tableWidth)
	}

	return &m, nil
}

func (m GameScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case TickMsg:

		m.lastUpdate = time.Now()
		updatedModel, err := m.refreshGameData()
		if err == nil {
			m = *updatedModel
		}

		cmds = append(cmds, tea.Every(REFRESH_GAME_RATE*time.Second, func(t time.Time) tea.Msg {
			return TickMsg(t)
		}))
		return m, tea.Batch(cmds...)

	case tea.WindowSizeMsg:
		constants.DocStyle = lipgloss.NewStyle().Width(msg.Width).Height(msg.Height).Padding(constants.VPadding, constants.HPadding)

		m.width = msg.Width - constants.DocStyle.GetHorizontalFrameSize()
		m.height = msg.Height - constants.DocStyle.GetVerticalFrameSize()

		helpTextHeight := lipgloss.Height(m.help.View(gameScreenKM))
		m.viewport.Width = m.width
		m.viewport.Height = m.height - helpTextHeight

		// Responsive Layout Logic
		var tableWidth int
		if m.width < RESPONSIVE_WIDTH_BREAKPOINT {
			tableWidth = m.width
		} else {
			tableWidth = m.width / 2
		}

		for i, playerTable := range m.playerTables {
			m.playerTables[i] = playerTable.WithTargetWidth(tableWidth)
		}

		m.linescoreTable = m.linescoreTable.WithTargetWidth(m.width)

		if m.popup != nil {
			m.popup = m.popup.Resize(msg, m.renderMainScreen())
		}

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, gameScreenKM.Quit):
			return m, tea.Quit
		case key.Matches(msg, gameScreenKM.Back):
			if m.popup == nil {
				return m.previousModel, tea.Batch(tea.ClearScreen)
			} else {
				gameScreenKM.SetEnabled(true)
				m.popup = nil
			}
		case key.Matches(msg, gameScreenKM.Up):
			m.viewport.LineUp(1)
		case key.Matches(msg, gameScreenKM.Down):
			m.viewport.LineDown(1)
		case key.Matches(msg, gameScreenKM.PageUp):
			m.viewport.ViewUp()
		case key.Matches(msg, gameScreenKM.PageDown):
			m.viewport.ViewDown()
		}
	}

	m.viewport.SetContent(m.renderMainScreen())
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m GameScreenModel) View() string {
	if m.popup != nil {
		return m.popup.View()
	}

	m.viewport.SetContent(m.renderMainScreen())

	helpContainer := lipgloss.NewStyle().
		SetString(m.help.View(gameScreenKM)).
		Width(m.width).
		Align(lipgloss.Left).
		String()

	ui := lipgloss.JoinVertical(lipgloss.Center, m.viewport.View(), helpContainer)

	return constants.DocStyle.Render(ui)
}

func (m GameScreenModel) renderMainScreen() string {

	updateTime := "Never"
	if !m.lastUpdate.IsZero() {
		updateTime = m.lastUpdate.Format("15:04:05")
	}

	scoreBox := components.RenderScoreText(&m.boxscore)

	// Responsive Layout Logic
	var playerStatsContent string
	awayTable := m.playerTables[TABLE_TO_INDEX_MAP["awayBatters"]].View()
	homeTable := m.playerTables[TABLE_TO_INDEX_MAP["homeBatters"]].View()

	if m.width < RESPONSIVE_WIDTH_BREAKPOINT {
		playerStatsContent = lipgloss.JoinVertical(lipgloss.Left, awayTable, homeTable)
	} else {
		playerStatsContent = lipgloss.JoinHorizontal(lipgloss.Top, awayTable, homeTable)
	}

	content := lipgloss.JoinVertical(lipgloss.Center, scoreBox, fmt.Sprintf("Update time: %s", updateTime), m.linescoreTable.View(), playerStatsContent)

	return content
}

func InitGameScreenModel(game ncaa.GameNCAA, previousModel Model) *GameScreenModel {
	boxscore, err := repositories.NewGameinfoRepository().GetGameinfoFromGameID(game.GameID)
	if err != nil {
		fmt.Printf("Failed to get boxscore")
		panic(err)
	}

	var (
		linescore      *ncaa.Linescore
		linescoreTable table.Model
		playerTables   []table.Model
	)

	if ls, err := repositories.NewLinescoreRepository().GetLinescoreFromGameId(game.GameID); err == nil {
		linescore = ls

		homeBattersTable := components.BuildPlayerStatsTable(linescore.TeamBoxscores[0].PlayerStats)
		awayBattersTable := components.BuildPlayerStatsTable(linescore.TeamBoxscores[1].PlayerStats)
		playerTables = []table.Model{awayBattersTable, homeBattersTable}

		linescoreTable = components.BuildLinescoreTable(game.Away.Names.Short, game.Home.Names.Short, linescore)
	} else {
		fmt.Printf("Linescore unavailable: %v\n", err)
		playerTables = []table.Model{components.EmptyPlayerStatsTable(), components.EmptyPlayerStatsTable()}
		linescoreTable = components.EmptyLinescoreTable()
	}

	vp := viewport.New(constants.WindowSize.Width, constants.WindowSize.Height)

	gameScreenModel := GameScreenModel{
		game:           game,
		previousModel:  previousModel,
		linescoreTable: linescoreTable,
		playerTables:   playerTables,
		boxscore:       *boxscore,
		help:           help.New(),
		viewport:       vp,
		width:          constants.WindowSize.Width,  // Set initial width
		height:         constants.WindowSize.Height, // Set initial height
	}

	return &gameScreenModel
}
