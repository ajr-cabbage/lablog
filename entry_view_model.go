package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type EntryViewModel struct {
	entry Entry
}

// Implement tea.Model interface
func (m *EntryViewModel) Init() tea.Cmd {
	return nil
}

func (m *EntryViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m EntryViewModel) View() string {
	return "loading..."
}

func NewEntryViewModel() *EntryViewModel {
	return &EntryViewModel{}
}
