package main

import (
	"database/sql"
	"log"
	"os"
	"path"

	"github.com/ajr-cabbage/lablog/internal/database"
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
	dbQueries := database.New(db)
	// apply up migration
	err = goose.SetDialect("sqlite")
	if err != nil {
		log.Fatal(err)
	}
	err = goose.Up(db, "sql/schema")
	if err != nil {
		log.Fatal(err)
	}

	m := NewMainModel(dbQueries)
	p := tea.NewProgram(m)
	_, err = p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
