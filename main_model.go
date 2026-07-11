package main

import (
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
	quitting     bool
}

func NewMainModel() *MainModel {
	var m MainModel
	m.listViewMod = NewListViewModel()
	lvMod, _ := m.listViewMod.(*ListViewModel)
	m.entryViewMod = NewEntryViewModel(lvMod)
	m.state = listView
	return &m
}

// Implement tea.Model interface
// TODO: define individual model inits before batching.
func (m *MainModel) Init() tea.Cmd {
	return nil
}

func (m *MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	// global key events
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "esc":
			if m.state == entryView {
				m.state = listView
			}
			return m, nil
		case "enter":
			if m.state == listView {
				m.state = entryView
			}
		}
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
	if m.quitting {
		return ""
	}
	switch m.state {
	case entryView:
		return m.entryViewMod.View()
	default:
		return m.listViewMod.View()
	}
}
