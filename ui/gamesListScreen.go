package ui

import (
	"fmt"
	"os"
	"time"

	"github.com/darekbienkowski/goncaa/api/ncaa"
	"github.com/darekbienkowski/goncaa/api/ncaa/repositories"
	"github.com/darekbienkowski/goncaa/ui/constants"
	"github.com/darekbienkowski/goncaa/ui/popup"
	"github.com/darekbienkowski/goncaa/ui/popup/dateJumpPopup"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const REFRESH_LIST_RATE = 30

type Model struct {
	date          time.Time
	gameList      list.Model
	help          help.Model
	jumpPopup     popup.IPopup
	width, height int
	lastUpdate    time.Time // New feat
}

var gamesListKM = GamesListKM{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Previous: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("left", "previous day"),
	),
	Next: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("right", "next day"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
	Jump: key.NewBinding(
		key.WithKeys("J"),
		key.WithHelp("J", "jump to date"),
	),
}

type TickMsg time.Time

func (m Model) Init() tea.Cmd {
	return tea.Every(REFRESH_LIST_RATE*time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case TickMsg:

		m.lastUpdate = time.Now()
		m = m.UpdateWithDate(m.date)

		cmds = append(cmds, tea.Every(REFRESH_LIST_RATE*time.Second, func(t time.Time) tea.Msg {
			return TickMsg(t)
		}))

		return m, tea.Batch(cmds...)
	case tea.WindowSizeMsg:
		constants.WindowSize = msg
		constants.DocStyle = lipgloss.NewStyle().Width(msg.Width).Height(msg.Height).Padding(constants.VPadding, constants.HPadding)

		m.width = msg.Width - constants.DocStyle.GetHorizontalFrameSize()
		m.height = msg.Height - constants.DocStyle.GetVerticalFrameSize()

		m.gameList.SetSize(m.width, m.height)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, gamesListKM.Quit):
			return m, tea.Quit
		case key.Matches(msg, gamesListKM.Next):
			// Update the model for tomorrow
			m = m.UpdateWithDate(m.date.AddDate(0, 0, 1))
		case key.Matches(msg, gamesListKM.Previous):
			//Update the model for yesterday
			m = m.UpdateWithDate(m.date.AddDate(0, 0, -1))
		case key.Matches(msg, gamesListKM.Enter):
			activeGame := m.gameList.SelectedItem().(ncaa.GameNCAA)
			gameScreenModel := InitGameScreenModel(activeGame, m)
			initCmd := gameScreenModel.Init()
			updateModel, updateCmd := gameScreenModel.Update(constants.WindowSize)
			return updateModel, tea.Batch(initCmd, updateCmd, tea.ClearScreen)
		case key.Matches(msg, gamesListKM.Jump):
			date := dateJumpDialog(m)
			m.jumpPopup = date
		}
	}

	m.gameList, cmd = m.gameList.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Initializing..."
	}

	if m.jumpPopup != nil {
		return m.jumpPopup.View()
	}

	return constants.DocStyle.Render(m.gameList.View())
}

func (m Model) UpdateWithDate(date time.Time) Model {
	schedule, err := repositories.NewScheduleRepository().GetScheduleForDate(date)
	if err != nil {
		panic(err)
	}

	m.date = date
	updateTime := "Never"
	if !m.lastUpdate.IsZero() {
		updateTime = m.lastUpdate.Format("15:04:05")
	}

	if len(schedule.Games) == 0 {
		m.gameList.SetItems([]list.Item{})
		m.gameList.Title = fmt.Sprintf("No games on %s, Update time: %s", date.Format(time.DateOnly), updateTime)
	} else {
		newListItems := gamesToItems(schedule.Games)
		m.gameList.SetItems(newListItems)
		m.gameList.Title = fmt.Sprintf("Games on %s, Update time: %s", date.Format(time.DateOnly), updateTime)
	}

	return m
}

func InitModel(date time.Time) tea.Model {
	schedule, err := repositories.NewScheduleRepository().GetScheduleForDate(date)

	if err != nil {
		panic(err)
	}

	if len(schedule.Games) == 0 {
		fmt.Println("No games on", date)
		os.Exit(0)
	}

	games := schedule.Games
	items := gamesToItems(games)

	// Custom styling
	customDelegate := list.NewDefaultDelegate()
	customDelegate.ShortHelpFunc = func() []key.Binding {
		return gamesListKM.ShortHelp()
	}

	customDelegate.FullHelpFunc = func() [][]key.Binding {
		return gamesListKM.FullHelp()
	}

	customDelegate.Styles.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(constants.PrimaryColor).
		Foreground(constants.PrimaryColor).
		Padding(0, 0, 0, 1)

	customDelegate.Styles.SelectedDesc = customDelegate.Styles.SelectedTitle

	m := Model{gameList: list.New(items, customDelegate, 0, 0), date: date, help: help.New()}
	updateTime := "Never"
	if !m.lastUpdate.IsZero() {
		updateTime = m.lastUpdate.Format("15:04:05")
	}
	m.gameList.Title = fmt.Sprintf("Games on %s, Update time: %s", date.Format(time.DateOnly), updateTime)

	return m
}

func gamesToItems(games []ncaa.GameWrapper) []list.Item {
	items := make([]list.Item, len(games))

	for i, game := range games {
		items[i] = list.Item(game.Game)
	}

	return items
}

func dateJumpDialog(m Model) dateJumpPopup.Model {
	// return current time in local time rounded to today's date
	popUp := dateJumpPopup.New(m.View(), m.width-2*constants.PopupHPadding, m.height-2*constants.PopupVPadding)
	return popUp
}
