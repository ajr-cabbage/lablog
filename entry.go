package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
)

type Entry struct {
	friendlyName string
	hostName     string
	description string
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
// implement DefaultDelegate interface
type CustomDelegate struct {}

func (d CustomDelegate) Height() int {return 3}

func (d CustomDelegate) Spacing() int {return 1}

func (d CustomDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {return nil}

func (d CustomDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	
}


