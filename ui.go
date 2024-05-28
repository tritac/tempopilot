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

	dateView := canvas.NewRectangle(color.Gray{Y: 0x99})
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
