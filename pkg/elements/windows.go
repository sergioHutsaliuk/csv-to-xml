//go:build windows

package elements

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
)

var filters = []cfd.FileFilter{
	{
		DisplayName: "CSV Files (*.csv)",
		Pattern:     "*.csv",
	},
	{
		DisplayName: "XML Files (*.xml)",
		Pattern:     "*.xml",
	},
}

func sourceSelection(filePathName binding.String, _ fyne.Window) {
	path, err := cfdutil.ShowOpenFileDialog(cfd.DialogConfig{
		Title:                   "Open A File",
		Role:                    "Open source file",
		FileFilters:             filters,
		SelectedFileFilterIndex: 0,
		FileName:                "",
		DefaultExtension:        "csv",
	})

	if err != nil && err != cfd.ErrorCancelled {
		log.Println(err)
	}

	if err == nil {
		_ = filePathName.Set(path)
	}
}

func destinationSelection(filePathName binding.String, w fyne.Window) {
	path, err := cfdutil.ShowSaveFileDialog(cfd.DialogConfig{
		Title:                   "Save A File",
		Role:                    "Save result file",
		FileFilters:             filters,
		SelectedFileFilterIndex: 1,
		FileName:                "",
		DefaultExtension:        "xml",
	})

	if err != nil && err != cfd.ErrorCancelled {
		log.Println(err)
	}

	if err == nil {
		_ = filePathName.Set(path)
	}
}
