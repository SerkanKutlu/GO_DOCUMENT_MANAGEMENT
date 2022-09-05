package utils

import (
	customerror "documentService/customError"
	"documentService/model"
	"errors"
	"io"
	"mime/multipart"
	"os"
)

var validTypes = [4]string{"application/pdf", "image/png", "image/jpg", "image/jpeg"}

func getMimeType(file *multipart.FileHeader) string {
	return file.Header.Get("Content-Type")
}

func isValidType(file *multipart.FileHeader) error {
	mime := getMimeType(file)
	for _, valid := range validTypes {
		if mime == valid {
			return nil
		}
	}
	return errors.New(mime + " is not a valid type")
}
func CheckFilesTypes(files *[]*multipart.FileHeader) error {
	for _, file := range *files {
		if err := isValidType(file); err != nil {
			return err
		}
	}
	return nil

}
func CopyFile(file *multipart.FileHeader, entity *model.Document) *customerror.CustomError {
	src, err := file.Open()
	if err != nil {
		return customerror.NewError(err.Error(), 500)
	}
	dst, err := os.Create(entity.Path)
	if _, err = io.Copy(dst, src); err != nil {
		return customerror.NewError(err.Error(), 500)
	}

	//CONVERT IMAGE TO PDF
	if getMimeType(file) != "application/pdf" {
		if err := ConvertToPdfAndSave(entity.Path); err != nil {
			return err
		}
		//Delete the old pdf file
		if err := os.RemoveAll(entity.Path); err != nil {
			return customerror.NewError(err.Error(), 500)
		}

	}
	return nil
}
