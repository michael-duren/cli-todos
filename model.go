package main

import (
	"context"
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
	context context.Context
}

var quitKeys = key.NewBinding(
	key.WithKeys("q", "esc", "ctrl+c"),
	key.WithHelp("", "press q to quit"),
)

func initialModel(q *database.Queries, ctx context.Context) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{spinner: s, page: HomePage, queries: q, context: ctx}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}
