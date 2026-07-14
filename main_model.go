package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Possible Views
type ViewState int

const (
	listView ViewState = iota
	entryView
	formView
)

type MainModel struct {
	state        ViewState
	listViewMod  tea.Model
	entryViewMod tea.Model
	formViewMod  tea.Model
	quitting     bool
}

func NewMainModel() *MainModel {
	var m MainModel
	m.listViewMod = NewListViewModel()
	lvMod, _ := m.listViewMod.(*ListViewModel)
	m.entryViewMod = NewEntryViewModel(lvMod)
	m.formViewMod = NewFormViewModel()
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
		case "ctrl+c", "ctrl+q":
			m.quitting = true
			return m, tea.Quit
		case "q":
			if m.state != formView {
				m.quitting = true
				return m, tea.Quit
			}
		case "esc":
			if m.state != listView {
				m.state = listView
			}
			return m, nil
		case "enter":
			if m.state == listView {
				m.state = entryView
			}
		case "+":
			if m.state == listView {
				m.state = formView
				f, ok := m.formViewMod.(*FormViewModel)
				if ok {
					f.initForm()
				}
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
	case formView:
		newFormMod, newCmd := m.formViewMod.Update(msg)
		m.formViewMod = newFormMod
		cmd = newCmd
	}
	return m, cmd
}

func (m MainModel) View() string {
	// skip final render
	if m.quitting {
		return ""
	}
	switch m.state {
	case entryView:
		return m.entryViewMod.View()
	case formView:
		return m.formViewMod.View()
	default:
		return m.listViewMod.View()
	}
}
