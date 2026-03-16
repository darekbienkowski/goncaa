package ui

import "github.com/charmbracelet/bubbles/key"

type GamesListKM struct {
	Up          key.Binding
	Down        key.Binding
	Previous    key.Binding
	Next        key.Binding
	Quit        key.Binding
	Enter       key.Binding
	FocusPicker key.Binding
	Jump        key.Binding
}

func (k GamesListKM) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

func (k GamesListKM) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Previous, k.Next, k.Jump, k.Quit, k.Enter}
}

type GameScreenKM struct {
	Up       key.Binding
	Down     key.Binding
	Back     key.Binding
	Quit     key.Binding
	PageUp   key.Binding
	PageDown key.Binding
}

func (k GameScreenKM) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

func (k GameScreenKM) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.PageUp, k.PageDown, k.Back, k.Quit}
}

func (k *GameScreenKM) SetEnabled(enabled bool) {
	k.Up.SetEnabled(enabled)
	k.Down.SetEnabled(enabled)
	k.PageUp.SetEnabled(enabled)
	k.PageDown.SetEnabled(enabled)
}
