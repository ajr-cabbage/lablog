package main

import (
	"context"

	"github.com/ajr-cabbage/lablog/internal/database"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type FormType int

const (
	InsertForm FormType = iota
	EditForm
	DeleteForm
)

var (
	AddFormStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#1c9e80")).
			Padding(1, 2).
			Align(lipgloss.Center)
	EditFormStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#1d6b91")).
			Padding(1, 2).
			Align(lipgloss.Center)
	DeleteFormStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#ba3c3c")).
			Padding(1, 2).
			Align(lipgloss.Center)
)

type FormViewModel struct {
	form     *huh.Form
	formType FormType
	width    int
	height   int
	db       *database.Queries
}

func NewFormViewModel(db *database.Queries) *FormViewModel {
	newFormMod := FormViewModel{}
	newFormMod.db = db
	return &newFormMod
}

// implement tea.Model interface
func (f *FormViewModel) Init() tea.Cmd {
	return f.form.Init()
}

func (f *FormViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		f.width = msg.Width
		f.height = msg.Height
	}
	formMod, cmd := f.form.Update(msg)
	if huhForm, ok := formMod.(*huh.Form); ok {
		f.form = huhForm
	}
	return f, cmd
}

func (f FormViewModel) View() string {
	switch f.formType {
	case InsertForm:
		return lipgloss.Place(f.width, f.height, lipgloss.Center, lipgloss.Center, AddFormStyle.Render(f.form.View()))
	case EditForm:
		return EditFormStyle.Render(f.form.View())
	default:
		return DeleteFormStyle.Render(f.form.View())
	}
}

func (f *FormViewModel) initForm(m *MainModel, fType FormType) {
	f.formType = fType
	switch fType {
	case InsertForm:
		f.form = huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[category]().
					Key("category").
					Options(
						huh.NewOption("Server", servers),
						huh.NewOption("Network Hardware", networkHardware),
						huh.NewOption("User Device", userMachines),
						huh.NewOption("Audio Device", audioDevices),
					).
					Title("Choose a Category"),
				huh.NewInput().
					Title("Friendly Name").
					Prompt("> ").
					Key("friendlyName"),
				huh.NewInput().
					Title("Host Name").
					Prompt("> ").
					Key("hostName"),
				huh.NewInput().
					Title("Description").
					Prompt("> ").
					Key("description"),
				huh.NewInput().
					Title("IP Address").
					Prompt("> ").
					Key("ipAddress"),
			).Title("Add Entry"),
		)
	case DeleteForm:
		f.form = huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Key("delete").
					Title("Are you sure you want to remove this device?"),
			),
		)
	case EditForm:
		// get info from current selected entry from db
		currentListView, ok := m.listViewMod.(*ListViewModel)
		if !ok {
			m.state = listView
		}
		focusedListItem, ok := currentListView.lists[currentListView.focused].Items()[currentListView.lists[currentListView.focused].Index()].(Entry)
		if !ok {
			m.state = listView
		}
		dbEntry, err := m.db.GetEntryByID(context.Background(), int64(focusedListItem.id))
		if err != nil {
			m.state = listView
		}
		// defaults to populate edit form
		catDefault := category(dbEntry.Category)
		nameDefault := dbEntry.FriendlyName
		hostDefault := dbEntry.HostName
		descDefault := dbEntry.Description
		ipDefault := dbEntry.IpAddress

		f.form = huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[category]().
					Key("category").
					Options(
						huh.NewOption("Server", servers),
						huh.NewOption("Network Hardware", networkHardware),
						huh.NewOption("User Device", userMachines),
						huh.NewOption("Audio Device", audioDevices),
					).
					Title("Choose a Category").
					Value(&catDefault),
				huh.NewInput().
					Title("Friendly Name").
					Prompt("> ").
					Key("friendlyName").
					Value(&nameDefault),
				huh.NewInput().
					Title("Host Name").
					Prompt("> ").
					Key("hostName").
					Value(&hostDefault),
				huh.NewInput().
					Title("Description").
					Prompt("> ").
					Key("description").
					Value(&descDefault),
				huh.NewInput().
					Title("IP Address").
					Prompt("> ").
					Key("ipAddress").
					Value(&ipDefault),
			).Title("Edit Entry"),
		)
	}
}
