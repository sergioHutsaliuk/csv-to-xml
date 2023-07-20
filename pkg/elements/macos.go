//go:build darwin

package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
)

func sourceSelection(filePathName binding.String, w fyne.Window) {
	dialog.ShowFileOpen(func(r fyne.URIReadCloser, err error) {
		if r != nil {
			path := r.URI().Path()
			_ = filePathName.Set(path)
		}
	}, w)
}

func destinationSelection(filePathName binding.String, w fyne.Window) {
	dialog.ShowFileSave(func(r fyne.URIWriteCloser, err error) {
		if r != nil {
			path := r.URI().Path()
			_ = filePathName.Set(path)
		}
	}, w)
}
