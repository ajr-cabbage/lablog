package main

import (
	"database/sql"
	"log"
	"os"
	"path"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

func main() {
	// make/connect to db at the user home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dbPath := path.Join(homeDir, ".lablog.db")
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = goose.SetDialect("sqlite")
	if err != nil {
		log.Fatal(err)
	}
	err = goose.Up(db, "sql/schema")

	m := NewMainModel()
	p := tea.NewProgram(m)
	_, err = p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
