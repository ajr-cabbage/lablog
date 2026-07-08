package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Possible Views
type ViewState int

const (
	listView ViewState = iota
	entryView
)

type MainModel struct {
	state        ViewState
	listViewMod  tea.Model
	entryViewMod tea.Model
}

func NewMainModel() *MainModel {
	var m MainModel
	m.listViewMod = NewListViewModel()
	m.entryViewMod = NewEntryViewModel()
	m.state = listView
	return &m
}

// Implement tea.Model interface
// TODO: define individual model inits before batching.
func (m *MainModel) Init() tea.Cmd {
	/*
		return tea.Batch(
			m.listViewMod.Init(),
			m.entryViewMod.Init(),
		)
	*/
	return nil
}

// TODO: handle global events then switch events to child model update funcs
func (m *MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg {
	case tea.KeyCtrlC:
		os.Exit(0)
	case tea.KeyEsc:
		m.state = listView
	}
	// pass message to active sub model
	switch m.state {
	case listView:
		newListMod, newCmd := m.listViewMod.Update(msg)
		m.listViewMod = newListMod
		cmd = newCmd
	case entryView:
		newEntryMod, newCmd := m.entryViewMod.Update(msg)
		m.entryViewMod = newEntryMod
		cmd = newCmd
	}
	return m, cmd
}

func (m MainModel) View() string {
	switch m.state {
	case entryView:
		return m.entryViewMod.View()
	default:
		return m.listViewMod.View()
	}
}
