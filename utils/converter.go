package utils

import (
	"github.com/signintech/gopdf"
)

func ConvertToPdfAndSave(filePath string) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	if err := pdf.Image(filePath, 0, 0, gopdf.PageSizeA4); err != nil {
		return err
	}
	if err := pdf.WritePdf(filePath + ".pdf"); err != nil {
		return err
	}
	return nil
}
