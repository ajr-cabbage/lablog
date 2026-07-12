package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type FormViewModel struct {
	form *huh.Form
}

func NewFormViewModel() *FormViewModel {
	newFormMod := FormViewModel{}
	newFormMod.initForm()
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
	if f.form.State == huh.StateCompleted {
		return fmt.Sprintf("Entry Added: %s", f.form.GetString("friendlyName"))
	}
	return f.form.View()
}

func (f *FormViewModel) initForm() {
	f.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("category").
				Options(huh.NewOptions("Server", "Network Hardware", "User Device")...).
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
