package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func Destination(filePathName binding.String, w fyne.Window) *fyne.Container {
	fileSelectionEntry := widget.NewEntryWithData(filePathName)
	fileSelectionEntry.SetPlaceHolder("Path to generated file")

	fileSelectionButton := widget.NewButton("Save to file", func() {
		destinationSelection(filePathName, w)
	})

	selectionLabel := widget.NewLabel("Save to")
	selectionBox := container.NewAdaptiveGrid(2, fileSelectionEntry, fileSelectionButton)
	fileSelectionEntry.Resize(fyne.Size{Width: 350})

	return container.NewVBox(selectionLabel, selectionBox)
}
