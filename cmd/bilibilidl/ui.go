package main

import (
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func selectList(title string, rows []string) (int, error) {
	if err := ui.Init(); err != nil {
		return 0, err
	}
	defer ui.Close()

	l := widgets.NewList()
	l.Title = title
	l.Rows = rows
	l.TextStyle = ui.NewStyle(ui.ColorBlue)
	l.WrapText = false
	l.SetRect(0, 0, 25, 8)

	ui.Render(l)

	for e := range ui.PollEvents() {
		switch e.ID {
		case "q", "<C-c>":
			ui.Close()
			os.Exit(0)
		case "<Down>":
			l.ScrollDown()
		case "<Up>":
			l.ScrollUp()
		case "<Enter>":
			return l.SelectedRow, nil
		}

		ui.Render(l)
	}
	return 0, nil
}
