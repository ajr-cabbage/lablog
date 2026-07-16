package main

import (
	"github.com/ajr-cabbage/lablog/internal/database"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
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
	db           *database.Queries
}

func NewMainModel(db *database.Queries) *MainModel {
	var m MainModel
	m.listViewMod = NewListViewModel(db)
	lvMod, _ := m.listViewMod.(*ListViewModel)
	m.entryViewMod = NewEntryViewModel(lvMod)
	m.formViewMod = NewFormViewModel(db)
	m.state = listView
	m.db = db
	return &m
}

// Implement tea.Model interface
// TODO: define individual model inits before batching.
func (m *MainModel) Init() tea.Cmd {
	return nil
}

func (m *MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
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
					f.initForm(m, InsertForm)
				}
			}
		case "delete":
			if m.state == listView {
				m.state = formView
				f, ok := m.formViewMod.(*FormViewModel)
				if ok {
					f.initForm(m, DeleteForm)
				}
			}
		case "u":
			if m.state == listView {
				m.state = formView
				f, ok := m.formViewMod.(*FormViewModel)
				if ok {
					f.initForm(m, EditForm)
				}
			}
		}
	}
	// pass message to active sub model
	switch m.state {
	case listView:
		newListMod, newCmd := m.listViewMod.Update(msg)
		m.listViewMod = newListMod
		cmds = append(cmds, newCmd)
	case entryView:
		newEntryMod, newCmd := m.entryViewMod.Update(msg)
		m.entryViewMod = newEntryMod
		cmds = append(cmds, newCmd)
	case formView:
		newFormMod, newCmd := m.formViewMod.Update(msg)
		f, ok := newFormMod.(*FormViewModel)
		if ok {
			m.formViewMod = f
		}
		cmds = append(cmds, newCmd)
		if f.form.State == huh.StateCompleted {
			switch f.formType {
			case InsertForm:
				addEntryHandler(m, f)
			case DeleteForm:
				deleteEntryHandler(m, f)
			case EditForm:
				editEntryHandler(m, f)
			}
		}
	}
	return m, tea.Batch(cmds...)
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
