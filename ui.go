package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type TempoGUI struct {
	win         fyne.Window
	preferences fyne.Preferences
}

func (g *TempoGUI) MakeGUI() fyne.CanvasObject {

	dateView := g.makeMonthWorkdaysList()
	workLogStats := widget.NewLabel("work log stats one ")
	workLogDetails := canvas.NewRectangle(color.Gray{Y: 0xff})
	objs := []fyne.CanvasObject{dateView, workLogStats, workLogDetails}
	return container.New(newTempoPilotLayout(dateView, workLogStats, workLogDetails), objs...)
}

func (g *TempoGUI) WelcomeAndKeyDialog() {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter your name")
	nameEntry.SetText(g.preferences.String("name"))

	apiEntry := widget.NewEntry()
	apiEntry.SetPlaceHolder("Enter Tempo your API key")

	tempoUserId := widget.NewEntry()
	tempoUserId.SetPlaceHolder("Tempo User Id")

	// calendar := xWiget.NewCalendar(time.Now(), func(time time.Time) {
	// 	fmt.Println(time)
	// })

	// nameLabel := widget.NewLabel("Name")
	tempoUserIdField := container.NewPadded(tempoUserId)
	apiEntryField := container.NewPadded(apiEntry)
	userNameField := container.NewPadded(nameEntry)

	inputFields := container.NewVBox(userNameField, apiEntryField, tempoUserIdField)

	// form := &widget.Form{
	// 	Items: []*widget.FormItem{
	// 		{Widget: nameInput},
	// 	},
	// }
	dialog := dialog.NewCustomConfirm("Welcome to TempoPilot", "Submit", "Cancel", inputFields, func(submitted bool) {

		if !submitted {
			g.win.Close()
		}

		fmt.Println("Name: ", nameEntry.Text)
		fmt.Println("API Key: ", apiEntry.Text)
		fmt.Println("Tempo User Id: ", tempoUserId.Text)
		g.preferences.SetString("name", nameEntry.Text)
		g.preferences.SetString("api_key", apiEntry.Text)
		g.preferences.SetString("user_id", tempoUserId.Text)

	}, g.win)
	dialog.Resize(fyne.Size{Width: 300, Height: 200})
	dialog.Show()

}

func (g *TempoGUI) makeMonthWorkdaysList() *widget.List {

	list := createList([]int{1, 2, 3, 4, 5, 6, 7})
	return list
}
func createList(data []int) *widget.List {
	list := widget.NewList(
		// func that returns the number of items in the list
		func() int {
			return len(data)
		},
		// func that returns the component structure of the List Item
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		// func that is called for each item in the list and allows
		// you to show the content on the previously defined ui structure
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(string("test"))
		})

	return list
}
