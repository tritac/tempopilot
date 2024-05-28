package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type workDay struct {
	widget.BaseWidget
	isWorkingDay  bool
	date          time.Time
	weekday       time.Weekday
	app           *TempoPolit
	homeContainer *fyne.Container
	timeButton    *widget.Button
}

func newWorkday(pilot *TempoPolit, homec *fyne.Container) *workDay {
	workday := &workDay{date: time.Now(), isWorkingDay: false, weekday: time.Now().Weekday(), app: pilot, homeContainer: homec}

	workday.ExtendBaseWidget(workday)

	workday.timeButton = widget.NewButtonWithIcon("A", theme.MenuDropDownIcon(), func() {
		position := fyne.CurrentApp().Driver().AbsolutePositionForObject(workday.timeButton)
		position.Y += workday.timeButton.MinSize().Height
	})
	workday.timeButton.Alignment = widget.ButtonAlignLeading
	workday.timeButton.IconPlacement = widget.ButtonIconTrailingText
	workday.timeButton.Importance = widget.LowImportance

	return workday
}

func (w *workDay) CreateRenderer() fyne.WidgetRenderer {
	op := canvas.NewRectangle(color.NRGBA{0x00, 0x00, 0x00, 0x59})

	c := container.NewStack(op, w.timeButton)

	return widget.NewSimpleRenderer(c)

}
