package handler

import (
	"documentService/dto"
	"documentService/model"
	"documentService/utils"
	"mime/multipart"
)

func (ds *DocumentService) UploadMultipleFiles(form *multipart.Form) (*dto.UploadResultDto, error) {
	files := form.File["documents"]
	if err := utils.CheckFilesTypes(&files); err != nil {
		return nil, err
	}

	var uploadedDocuments []model.Document
	for _, file := range files {
		documentEntity := model.CreateDocumentEntity(file)
		if err := utils.CopyFile(file, documentEntity); err != nil {
			return nil, err
		}

		uploadedDocuments = append(uploadedDocuments, *documentEntity)
	}
	if err := ds.MongoService.InsertMany(&uploadedDocuments); err != nil {
		return nil, err
	}

	return dto.CreateUploadResultDto(&uploadedDocuments), nil
}
