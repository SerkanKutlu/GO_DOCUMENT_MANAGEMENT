package utils

import (
	"archive/zip"
	customerror "documentService/customError"
	"documentService/model"
	uuid "github.com/satori/go.uuid"
	"io"
	"os"
)

func CreateZip(paths *[]model.DownloadModel) (*string, *customerror.CustomError) {
	zipId := uuid.NewV4().String()
	wd, err := os.Getwd()
	if err != nil {
		return nil, customerror.NewError(err.Error(), 500)
	}
	zipPath := wd + "\\documents\\temp\\" + zipId + ".zip"
	archive, err := os.Create(zipPath)
	if err != nil {
		return nil, customerror.NewError(err.Error(), 500)
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	for _, path := range *paths {
		file, err := os.Open(path.Path)
		if err != nil {
			return nil, customerror.NewError(err.Error(), 500)
		}
		w1, err := zipWriter.Create(path.FileName)
		if err != nil {
			return nil, customerror.NewError(err.Error(), 500)
		}
		if _, err := io.Copy(w1, file); err != nil {
			panic(err)
		}
		file.Close()
	}
	zipWriter.Close()
	return &zipPath, nil
}
