package logger

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	Log *log.Logger
	f   *os.File
)

func Init() {
	var err error
	f, err = tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("err %s", err)
	}

	Log = log.New(f, "debug: ", log.LstdFlags)
}

func LogAndSync(msg string) {
	Log.Println(msg)
	_ = f.Sync()
}

func Close() {
	f.Close()
}
