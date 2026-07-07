package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type category int

const (
	servers category = iota
	networkHardware
	userMachines
)

type Model struct {
	lists   []list.Model
	focused category
	loaded  bool
}

func NewModel() *Model {
	return &Model{}
}

func (m *Model) initLists(width, height int) {

	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}
	m.lists[servers].Title = "Servers"
	m.lists[servers].SetItems([]list.Item{
		Entry{friendlyName: "NAS", hostName: "thatnas", ipAddress: "123.255.255.122", desccription: "stores the files", online: true},
		Entry{friendlyName: "App Server", hostName: "app-lord", ipAddress: "123.255.255.120", desccription: "runs the apps", online: true},
	})
	m.lists[networkHardware].Title = "Network Hardware"
	m.lists[networkHardware].SetItems([]list.Item{
		Entry{friendlyName: "Router/Gateway", hostName: "gateway", ipAddress: "192.168.1.0", desccription: "ISP router", online: true},
		Entry{friendlyName: "Switch", hostName: "ugreen-2.5g", ipAddress: "123.255.255.111", desccription: "High speed lan switch", online: true},
	})
	m.lists[userMachines].Title = "User Devices"
	m.lists[userMachines].SetItems([]list.Item{
		Entry{friendlyName: "Nice PC", hostName: "framework", ipAddress: "192.168.1.55", desccription: "very fast desktop", online: true},
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.initLists(msg.Width, msg.Height)
		if !m.loaded {
			m.loaded = true
		}
	}
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.loaded {
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.lists[servers].View(),
			m.lists[networkHardware].View(),
			m.lists[userMachines].View(),
		)
	} else {
		return "loading..."
	}
}
