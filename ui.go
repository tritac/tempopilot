package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/tritac/tempopilot/worklogs"
)

type TempoGUI struct {
	win         fyne.Window
	preferences fyne.Preferences
	data        *worklogs.WorkLogStore
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

	list := createList(g.data.GenerateMonthDays(2024, time.Now().Month()))
	return list
}
func createList(data []worklogs.WorkDay) *widget.List {

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return canvas.NewText("template", color.NRGBA{R: 0, G: 255, B: 0, A: 255})
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*canvas.Text).Text = data[i].WeekDay.String()
		})

	return list
}

func listItem(date string, weekDay string) fyne.CanvasObject {
	return container.NewBorder(
		nil, nil, nil,
		// left of the border
		widget.NewCheck("", func(b bool) {}),
		// takes the rest of the space
		widget.NewLabel(date),
	)
}

func dateRowList() fyne.CanvasObject {
	return container.NewGridWithColumns(3,
		widget.NewLabel("asas"),
		container.NewMax(),
		rightAligned(
			container.NewHBox(
				widget.NewButtonWithIcon("", theme.ZoomInIcon(), func() {}),
				widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {}),
			),
		),
	)
}

func rightAligned(object fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(nil, nil, nil, object)
}
