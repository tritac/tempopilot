package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/tritac/tempopilot/worklogs"
)

type TempoPolit struct {
	main  fyne.Window
	store *worklogs.WorkLogStore
}

func main() {
	a := app.NewWithID("com.tempopilot.amit")
	w := a.NewWindow("Tempo Pilot")
	gui := TempoGUI{win: w, preferences: a.Preferences()}
	gui.WelcomeAndKeyDialog()
	w.SetContent(container.NewStack(gui.MakeGUI()))
	w.ShowAndRun()
}

// func (n *TempoPolit) home() fyne.CanvasObject {
// 	layout := &tempoPilotLayout{}
// 	homeContainer := container.New(layout)

// 	// cells := []fyne.CanvasObject{}

// 	// homeContainer.Objects = append(cells, newWorkday(n, homeContainer))
// 	// scroll := container.NewVScroll(homeContainer)

// 	// scroll.SetMinSize(layout.minOuterSize())
// 	// return scroll
// 	// text1 := canvas.NewText("1", color.White)
// 	// text2 := canvas.NewText("2", color.White)
// 	// text3 := canvas.NewText("3", color.White)
// 	// grid := container.New(layout.NewGridLayout(2), text1, text2, text3)
// 	// return grid
// 	return homeContainer

// }
