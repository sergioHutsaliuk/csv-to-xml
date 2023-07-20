package convertor

import (
	"errors"
	"os"

	"fyne.io/fyne/v2/data/binding"
	"github.com/gocarina/gocsv"
)

func Convert(source binding.String, destination binding.String) error {
	sourcePath, err := source.Get()
	switch {
	case err != nil:
		return err
	case sourcePath == "":
		return errors.New("source file not selected or doesn't exist")
	}

	destinationPath, err := destination.Get()
	switch {
	case err != nil:
		return err
	case destinationPath == "":
		return errors.New("destination file not selected or doesn't exist")
	}

	err = process(sourcePath, destinationPath)
	if err != nil {
		return err
	}

	return nil
}

func process(sourcePath, destinationPath string) error {
	var records []*DataCSV

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}

	defer func(file ...*os.File) {
		for _, f := range file {
			_ = f.Close()
		}
	}(sourceFile, destinationFile)

	if err = gocsv.UnmarshalFile(sourceFile, &records); err != nil {
		panic(err)
	}

	_, err = destinationFile.WriteString("<Facturi>\n")

	for _, v := range records {
		v.preRender()

		switch v.VAT {
		case "0.00%":
			_, err = destinationFile.WriteString(withoutVAT(v))
		default:
			_, err = destinationFile.WriteString(withVAT(v))
		}
	}

	_, err = destinationFile.WriteString("</Facturi>")

	return nil
}
