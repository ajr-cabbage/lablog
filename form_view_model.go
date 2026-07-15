package main

import (
	"github.com/ajr-cabbage/lablog/internal/database"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type FormViewModel struct {
	form *huh.Form
	db   *database.Queries
}

func NewFormViewModel(db *database.Queries) *FormViewModel {
	newFormMod := FormViewModel{}
	newFormMod.initForm()
	newFormMod.db = db
	return &newFormMod
}

// implement tea.Model interface
func (f *FormViewModel) Init() tea.Cmd {
	return f.form.Init()
}

func (f *FormViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	formMod, cmd := f.form.Update(msg)
	if huhForm, ok := formMod.(*huh.Form); ok {
		f.form = huhForm
	}
	return f, cmd
}

func (f FormViewModel) View() string {
	return f.form.View()
}

func (f *FormViewModel) initForm() {
	f.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[category]().
				Key("category").
				Options(huh.NewOption("Server", servers),
					huh.NewOption("Network Hardware", networkHardware),
					huh.NewOption("User Device", userMachines),
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
		),
	)
}
