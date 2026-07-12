package main

import (
	"github.com/ajr-cabbage/lablog/tui"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	m := tui.NewMainModel()
	p := tea.NewProgram(m)
	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
