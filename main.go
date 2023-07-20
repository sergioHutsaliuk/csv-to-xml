package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/z1od/csv-to-xml/pkg/elements"
)

func main() {
	a := app.New()

	w := a.NewWindow("CSVtoXML")
	w.Resize(fyne.Size{Width: 400})
	w.SetFixedSize(true)

	sourcePathName := binding.NewString()
	destinationPathName := binding.NewString()
	text := binding.NewString()

	selection := elements.Selection(sourcePathName, w)
	saving := elements.Destination(destinationPathName, w)
	textSection := widget.NewLabelWithData(text)
	button := elements.Button(sourcePathName, destinationPathName, text)

	content := container.NewVBox(
		selection,
		widget.NewSeparator(),
		saving,
		widget.NewSeparator(),
		textSection,
		widget.NewSeparator(),
		button,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
