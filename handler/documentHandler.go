package handler

import (
	"documentService/dto"
	"documentService/model"
	"documentService/utils"
	"fmt"
	"mime/multipart"
)

func (ds *DocumentService) UploadMultipleFiles(form *multipart.Form) (*dto.UploadResultDto, error) {
	files := form.File["documents"]
	if err := utils.CheckFilesTypes(&files); err != nil {
		return nil, err
	}
	uploadResults := new(dto.UploadResultDto)             //Return pointer
	uploadResults.UploadedFiles = make(map[string]string) //Return objects itself
	for _, file := range files {
		if err := utils.CopyFile(file); err != nil {
			return nil, err
		}
		documentEntity := model.CreateDocumentEntity(file)
		fmt.Println(documentEntity)
	}
	return nil, nil
}
