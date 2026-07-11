package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type EntryViewModel struct {
	entry       Entry
	parentModel *ListViewModel
	loaded      bool
}

// Implement tea.Model interface
func (m *EntryViewModel) Init() tea.Cmd {
	return nil
}

func (m *EntryViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.initEntry()
	return m, nil
}

func (m EntryViewModel) View() string {
	return fmt.Sprint(m.entry)
}

func (m *EntryViewModel) initEntry() {
	focusedListItem := m.parentModel.lists[m.parentModel.focused].Items()[m.parentModel.lists[m.parentModel.focused].Index()]
	entryItem, ok := focusedListItem.(Entry)
	if !ok {
	}
	m.entry = entryItem
}

func NewEntryViewModel(l *ListViewModel) *EntryViewModel {
	return &EntryViewModel{
		parentModel: l,
	}
}
