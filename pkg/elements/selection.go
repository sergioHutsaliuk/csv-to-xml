package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func Selection(filePathName binding.String, w fyne.Window) *fyne.Container {
	fileSelectionEntry := widget.NewEntryWithData(filePathName)
	fileSelectionEntry.SetPlaceHolder("Path to source file")

	fileSelectionButton := widget.NewButton("Open file", func() {
		sourceSelection(filePathName, w)
	})

	selectionLabel := widget.NewLabel("Read from")
	selectionBox := container.NewAdaptiveGrid(2, fileSelectionEntry, fileSelectionButton)
	fileSelectionEntry.Resize(fyne.Size{Width: 350})

	return container.NewVBox(selectionLabel, selectionBox)
}
