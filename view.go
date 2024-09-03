package main

import (
	"d-todo/views"
	"fmt"
)

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	if m.width == 0 {
		str := fmt.Sprintf("\n\n   %s Loading forever... %s\n\n", m.spinner.View(), quitKeys.Help().Desc)
		if m.quitting {
			return str + "\n"
		}
		return str
	}

	// pages
	switch m.page {
	default:
		return views.HomePage()
	}
}
