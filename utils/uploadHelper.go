package utils

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
)

var validTypes = [3]string{"application/pdf", "image/png", "image/jpg"}

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

func CopyFile(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}

	dst, err := os.Create("documents/" + file.Filename)
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}
