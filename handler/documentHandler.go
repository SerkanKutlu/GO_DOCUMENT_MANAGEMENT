package handler

import (
	customerror "documentService/customError"
	"documentService/dto"
	"documentService/model"
	"documentService/utils"
	"github.com/golang-jwt/jwt"
	"mime/multipart"
	"os"
)

func (ds *DocumentService) UploadMultipleFiles(form *multipart.Form, authUser *jwt.Token) (*dto.UploadResultDto, *customerror.CustomError) {
	files := form.File["documents"]
	if err := utils.CheckFilesTypes(&files); err != nil {
		return nil, customerror.UnSupportedFileType
	}
	var uploadedDocuments []model.Document
	userId := utils.GetUserId(authUser)
	for _, file := range files {
		documentEntity := model.CreateDocumentEntity(file, userId)
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
func (ds *DocumentService) ShowAllEntity(authUser *jwt.Token) (*[]model.Document, *customerror.CustomError) {
	userRole := utils.GetUserRole(authUser)
	var documents *[]model.Document
	var err *customerror.CustomError
	if userRole == "User" {
		userId := utils.GetUserId(authUser)
		documents, err = ds.MongoService.GetOfUsers(userId)
	} else {
		documents, err = ds.MongoService.GetAll()
	}
	if err != nil {
		return nil, err
	}
	return documents, nil

}
func (ds *DocumentService) DeleteEntity(id string, authUser *jwt.Token) *customerror.CustomError {
	userRole := utils.GetUserRole(authUser)
	var path *string
	var err *customerror.CustomError
	if userRole == "User" {
		userId := utils.GetUserId(authUser)
		path, err = ds.MongoService.DeleteWithUserId(id, userId)
	} else {
		path, err = ds.MongoService.Delete(id)
	}
	if err != nil {
		return err
	}
	if err := os.RemoveAll(*path); err != nil {
		return customerror.NewError("File not deleted from disk", 500)
	}
	return nil
}
