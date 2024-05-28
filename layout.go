// package main

// import (
// 	"math"

// 	"fyne.io/fyne/v2"
// )

// const (
// 	minWidth  float32 = 600
// 	minHeight float32 = 400
// 	cellSpace float32 = 16
// )

// type tempoPilotLayout struct {
// 	cols int
// }

// func (l *tempoPilotLayout) Layout(objs []fyne.CanvasObject, s fyne.Size) {
// 	l.cols = 10
// }

// func (l *tempoPilotLayout) MinSize(cells []fyne.CanvasObject) fyne.Size {
// 	cols := l.cols
// 	if cols < 1 { // possibly not layed out yet
// 		cols = 1
// 	}
// 	rows := int(math.Ceil(float64(len(cells)) / float64(l.cols)))
// 	height := float32((minHeight+cellSpace*4)*float32(rows) - cellSpace)
// 	if fyne.CurrentDevice().IsMobile() {
// 		return fyne.NewSize(minWidth, height)
// 	}

// 	return fyne.NewSize(minWidth, height+cellSpace*2)
// }

// func (l *tempoPilotLayout) minOuterSize() fyne.Size {
// 	if fyne.CurrentDevice().IsMobile() {
// 		return fyne.NewSize(minWidth, minHeight)
// 	}
// 	return fyne.NewSize(minWidth+cellSpace*2, minHeight+cellSpace*2)

// }

package main

import "fyne.io/fyne/v2"

type TempoPilotLayout struct {
	dateView, workLogStats, workLogDetails fyne.CanvasObject
}

const (
	sizeWidth = 200
)

func newTempoPilotLayout(dateView, workLogStats, workLogDetails fyne.CanvasObject) *TempoPilotLayout {
	return &TempoPilotLayout{dateView: dateView, workLogStats: workLogStats, workLogDetails: workLogDetails}
}

func (l *TempoPilotLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	screen := size
	topHeight := 50.0

	l.dateView.Resize(fyne.NewSize(sizeWidth, screen.Height-float32(topHeight)))
	l.dateView.Move(fyne.NewPos(0, float32(topHeight)))

	l.workLogDetails.Resize(fyne.NewSize(screen.Width-sizeWidth, screen.Height-float32(topHeight)))
	l.workLogDetails.Move(fyne.NewPos(sizeWidth, float32(topHeight)))

	l.workLogStats.Resize(fyne.NewSize(screen.Width-sizeWidth, float32(topHeight)))
	l.workLogStats.Move(fyne.NewPos(sizeWidth, 0))

}

func (l *TempoPilotLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(800, 600)
}
