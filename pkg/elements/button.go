package elements

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	"github.com/z1od/csv-to-xml/pkg/convertor"
)

func Button(source binding.String, destination binding.String, text binding.String) *widget.Button {
	return widget.NewButton("Start conversion", func() {
		err := convertor.Convert(source, destination)
		switch err {
		case nil:
			_ = text.Set("Done!")
		default:
			_ = text.Set("Error: " + err.Error())
		}
	})
}
