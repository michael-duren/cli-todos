package main

import (
	"context"
	"d-todo/database"
	"d-todo/logger"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		fmt.Println("\nReceived interrupt signal, shutting down")
		cancel()
	}()

	db, err := sql.Open("sqlite3", "database.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logger.Init()
	defer logger.Close()

	q := database.New(db)
	p := tea.NewProgram(initialModel(q), tea.WithAltScreen())

	select {
	case <-ctx.Done():
		fmt.Println("Context canceled, exiting")
	default:
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}
	}

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
