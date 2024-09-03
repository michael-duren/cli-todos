package main

import (
	"d-todo/database"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Page int

const (
	HomePage         Page = iota
	AddTodoPage           = iota
	AddDailyTodoPage      = iota
)

type errMsg error

type model struct {
	err      error
	spinner  spinner.Model
	quitting bool
	page     Page
	width    int
	height   int

	queries *database.Queries
}

var quitKeys = key.NewBinding(
	key.WithKeys("q", "esc", "ctrl+c"),
	key.WithHelp("", "press q to quit"),
)

func initialModel(q *database.Queries) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{spinner: s, page: HomePage, queries: q}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}
