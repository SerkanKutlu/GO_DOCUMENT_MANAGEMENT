package dto

import "documentService/model"

type UploadResultDto struct {
	UploadedFiles map[string]string
}

func CreateUploadResultDto(uploadedFiles *[]model.Document) *UploadResultDto {
	uploadResults := new(UploadResultDto)                 //Return pointer
	uploadResults.UploadedFiles = make(map[string]string) //Return objects itself
	for _, file := range *uploadedFiles {
		uploadResults.UploadedFiles[file.FileName] = file.Id
	}
	return uploadResults
}
