package utils

import (
	customerror "documentService/customError"
	"github.com/signintech/gopdf"
)

func ConvertToPdfAndSave(filePath string) *customerror.CustomError {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	if err := pdf.Image(filePath, 0, 0, gopdf.PageSizeA4); err != nil {
		return customerror.NewError(err.Error(), 500)
	}
	if err := pdf.WritePdf(filePath + ".pdf"); err != nil {
		return customerror.NewError(err.Error(), 500)
	}
	return nil
}
