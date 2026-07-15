package main

import (
	"context"
	"fmt"

	"github.com/ajr-cabbage/lablog/internal/database"
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
	db      *database.Queries
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
			l.lists[l.focused].SetDelegate(NewCustomDelegate(false))
			if l.focused == userMachines {
				l.focused = servers
			} else {
				l.focused++
			}
			l.lists[l.focused].SetDelegate(NewCustomDelegate(true))
		case "left":
			l.lists[l.focused].SetDelegate(NewCustomDelegate(false))
			if l.focused == servers {
				l.focused = userMachines
			} else {
				l.focused--
			}
			l.lists[l.focused].SetDelegate(NewCustomDelegate(true))
		case "r": // refresh whenever returning to list view
			return l, l.refreshList()
		}
	}
	var cmd tea.Cmd
	l.lists[l.focused], cmd = l.lists[l.focused].Update(msg)
	return l, cmd
}

// combine list models
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

func NewListViewModel(db *database.Queries) *ListViewModel {
	return &ListViewModel{db: db}
}

// Override default delegate styles and provide alt styles for unfocused lists
func NewCustomDelegate(focused bool) list.DefaultDelegate {
	// define styles per condition
	listFocusedItemStyles := list.NewDefaultItemStyles()
	listNotFocusedItemStyles := list.DefaultItemStyles{
		NormalTitle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#dddddd")).Padding(0, 0, 0, 2),
		SelectedTitle: lipgloss.NewStyle().Foreground(lipgloss.Color("#dddddd")).Padding(0, 0, 0, 2),
		NormalDesc:    lipgloss.NewStyle().Foreground(lipgloss.Color("#A49FA5")).Padding(0, 0, 0, 2),
		SelectedDesc:  lipgloss.NewStyle().Foreground(lipgloss.Color("#A49FA5")).Padding(0, 0, 0, 2),
	}
	// init delegate
	customDelegate := list.NewDefaultDelegate()
	customDelegate.SetHeight(3)
	// apply correct style
	if focused {
		customDelegate.Styles = listFocusedItemStyles
	} else {
		customDelegate.Styles = listNotFocusedItemStyles
	}
	return customDelegate
}

// initialize styles and refresh() data from database
func (l *ListViewModel) initLists(width, height int) {
	listWidth := width/3 - 2
	listWidth = max(listWidth, 10)
	focusedDelegate := NewCustomDelegate(true)
	notFocusedDelegate := NewCustomDelegate(false)
	// d.Styles.SelectedDesc = d.Styles.NormalDesc
	serversList := list.New([]list.Item{}, focusedDelegate, listWidth, height)
	serversList.SetShowHelp(false)
	networkList := list.New([]list.Item{}, notFocusedDelegate, listWidth, height)
	networkList.SetShowHelp(false)
	userList := list.New([]list.Item{}, notFocusedDelegate, listWidth, height)
	userList.SetShowHelp(false)
	l.lists = []list.Model{serversList, networkList, userList}
	l.refreshList()
}

func (l *ListViewModel) refreshList() tea.Cmd {
	categories := []category{servers, networkHardware, userMachines}
	var cmds []tea.Cmd
	for _, cat := range categories {
		switch cat {
		case servers:
			l.lists[cat].Title = "Servers"
		case networkHardware:
			l.lists[cat].Title = "Network Hardware"
		case userMachines:
			l.lists[cat].Title = "User Devices"
		}

		dbEntries, err := l.db.GetEntriesByCategory(context.Background(), int64(cat))
		if err != nil {
			fmt.Println(err)
		}
		var newItems []list.Item
		for _, entry := range dbEntries {
			//TODO: tie online field to ping result
			rawListEntry := Entry{
				id:           int(entry.ID),
				friendlyName: entry.FriendlyName,
				hostName:     entry.HostName,
				description:  entry.Description,
				ipAddress:    entry.IpAddress,
				online:       true,
			}
			newItems = append(newItems, rawListEntry)
		}
		cmd := l.lists[cat].SetItems(newItems)
		cmds = append(cmds, cmd)
	}

	return tea.Batch(cmds...)
}
