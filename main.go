package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := NewMainModel()
	p := tea.NewProgram(m)
	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
