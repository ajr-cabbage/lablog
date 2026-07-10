package main

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Entry struct {
	friendlyName string
	hostName     string
	description  string
	ipAddress    string
	online       bool
}

// implement list.Item interface
func (e Entry) FilterValue() string {
	return e.friendlyName + " " + e.hostName + e.ipAddress
}

func (e Entry) Title() string {
	return e.friendlyName
}

func (e Entry) Description() string {
	return e.description
}

// implement list.ItemDelegate interface
type CustomDelegate struct{}

func (d CustomDelegate) Height() int { return 3 }

func (d CustomDelegate) Spacing() int { return 1 }

func (d CustomDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d CustomDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(Entry)
	if !ok {
		fmt.Println("Can't assert list item to Entry{}")
		return
	}

	isSelected := index == m.Index()
	selectedStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("205")).
		Padding(0, 1)
	notSelectedStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("241")).
		Padding(0, 1)
	onlineStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00ff00"))
	offlineStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ff0000"))

	var output string
	if item.online {
		output = fmt.Sprintf(
			"%s\n%s\n%s\n",
			item.friendlyName,
			item.ipAddress,
			onlineStyle.Render("Online"),
		)
	} else {
		output = fmt.Sprintf(
			"%s\n%s\n%s\n",
			item.friendlyName,
			item.ipAddress,
			offlineStyle.Render("Offline"),
		)
	}

	if isSelected {
		fmt.Fprint(w, selectedStyle.Render(output))
	} else {
		fmt.Fprint(w, notSelectedStyle.Render(output))
	}

}
