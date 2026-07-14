package main

import (
	"fmt"

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
	return e.friendlyName + " " + e.ipAddress
}

func (e Entry) Title() string {
	return e.friendlyName
}

func (e Entry) Description() string {
	var output string
	// custom output style
	if e.online {
		output = fmt.Sprintf(
			"%s\n%s",
			lipgloss.NewStyle().Foreground(lipgloss.Color("#bbbbbb")).Render(e.ipAddress),
			lipgloss.NewStyle().Foreground(lipgloss.Color("#8fce00")).Render("Online"),
		)
	} else {
		output = fmt.Sprintf(
			"%s\n%s",
			lipgloss.NewStyle().Foreground(lipgloss.Color("#bbbbbb")).Render(e.ipAddress),
			lipgloss.NewStyle().Foreground(lipgloss.Color("#ba3c3c")).Render("Offline"),
		)
	}
	return output
}
