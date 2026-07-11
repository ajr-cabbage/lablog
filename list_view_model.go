package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Possible List Categories
type category int

const (
	servers category = iota
	networkHardware
	userMachines
)

type ListViewModel struct {
	lists   []list.Model
	focused category
	loaded  bool
}

// Implement tea.Model interface
// TODO: fetch database info
func (l *ListViewModel) Init() tea.Cmd {
	return nil
}

func (l *ListViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !l.loaded {
			l.initLists(msg.Width, msg.Height)
			l.loaded = true
		} else {
			for i := range l.lists {
				l.lists[i].SetWidth(msg.Width/3 - 2)
				l.lists[i].SetHeight(msg.Height)
			}
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "right":
			if l.focused == userMachines {
				l.focused = servers
			} else {
				l.focused++
			}
		case "left":
			if l.focused == servers {
				l.focused = userMachines
			} else {
				l.focused--
			}
		}
	}

	var cmd tea.Cmd
	l.lists[l.focused], cmd = l.lists[l.focused].Update(msg)
	return l, cmd
}

func (l ListViewModel) View() string {
	if l.loaded {
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			l.lists[servers].View(), " ",
			l.lists[networkHardware].View(), " ",
			l.lists[userMachines].View(),
		)
	} else {
		return "loading..."
	}
}

func NewListViewModel() *ListViewModel {
	return &ListViewModel{}
}

// dummy initial data for testing
func (l *ListViewModel) initLists(width, height int) {
	listWidth := width/3 - 2
	listWidth = max(listWidth, 10)
	d := list.NewDefaultDelegate()
	d.SetHeight(3)
	// d.Styles.SelectedDesc = d.Styles.NormalDesc
	defaultList := list.New([]list.Item{}, d, listWidth, height)
	defaultList.SetShowHelp(false)
	l.lists = []list.Model{defaultList, defaultList, defaultList}
	l.lists[servers].Title = "Servers"
	l.lists[servers].SetItems([]list.Item{
		Entry{friendlyName: "NAS", hostName: "thatnas", ipAddress: "123.255.255.122", description: "stores the files", online: true},
		Entry{friendlyName: "App Server", hostName: "app-lord", ipAddress: "123.255.255.120", description: "runs the apps", online: false},
		Entry{friendlyName: "App Server", hostName: "app-lord", ipAddress: "123.255.255.120", description: "runs the apps", online: false},
		Entry{friendlyName: "App Server", hostName: "app-lord", ipAddress: "123.255.255.120", description: "runs the apps", online: false},
		Entry{friendlyName: "App Server", hostName: "app-lord", ipAddress: "123.255.255.120", description: "runs the apps", online: false},
		Entry{friendlyName: "App Server", hostName: "app-lord", ipAddress: "123.255.255.120", description: "runs the apps", online: false},
		Entry{friendlyName: "NAS", hostName: "thatnas", ipAddress: "123.255.255.122", description: "stores the files", online: true},
		Entry{friendlyName: "App Server", hostName: "app-lord", ipAddress: "123.255.255.120", description: "runs the apps", online: false},
		Entry{friendlyName: "App Server", hostName: "app-lord", ipAddress: "123.255.255.120", description: "runs the apps", online: false},
		Entry{friendlyName: "App Server", hostName: "app-lord", ipAddress: "123.255.255.120", description: "runs the apps", online: false},
	})
	l.lists[networkHardware].Title = "Network Hardware"
	l.lists[networkHardware].SetItems([]list.Item{
		Entry{friendlyName: "Router/Gateway", hostName: "gateway", ipAddress: "192.168.1.0", description: "ISP router", online: true},
		Entry{friendlyName: "Switch", hostName: "ugreen-2.5g", ipAddress: "123.255.255.111", description: "High speed lan switch", online: true},
	})
	l.lists[userMachines].Title = "User Devices"
	l.lists[userMachines].SetItems([]list.Item{
		Entry{friendlyName: "Nice PC", hostName: "framework", ipAddress: "192.168.1.55", description: "very fast desktop", online: true},
	})
}
