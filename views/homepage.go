package views

import (
	"d-todo/database"
	"d-todo/logger"
	"fmt"
)

type HomePageModel struct {
	Todos *[]database.Todo
}

func HomePage(props *HomePageModel) string {
	title := "Home Page\n\n"
	logger.LogAndSync(fmt.Sprintf("todos: %d", len(*props.Todos)))
	var todoStr string
	for _, todo := range *props.Todos {
		todoStr += fmt.Sprintf("Todo: %s\n\n", todo.Name)
	}
	return title + todoStr
}
