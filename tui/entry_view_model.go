package tui

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
	// TODO: only init when the entry view when it becomes focused
	m.initEntry()
	return m, nil
}

// TODO: Consider viewing more bubble elements with the struct info
func (m EntryViewModel) View() string {
	return fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		m.entry.friendlyName,
		m.entry.hostName,
		m.entry.description,
		m.entry.ipAddress,
	)
}

// find the focused item at the time of the update
func (m *EntryViewModel) initEntry() {
	focusedListItem := m.parentModel.
		lists[m.parentModel.focused].
		Items()[m.parentModel.lists[m.parentModel.focused].Index()]
	entryItem, ok := focusedListItem.(Entry)
	if ok {
		m.entry = entryItem
	}
}

// passes pointer to the focused list
func NewEntryViewModel(l *ListViewModel) *EntryViewModel {
	return &EntryViewModel{
		parentModel: l,
	}
}
