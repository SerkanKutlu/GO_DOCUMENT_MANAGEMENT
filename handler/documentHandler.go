package handler

import (
	customerror "documentService/customError"
	"documentService/dto"
	"documentService/model"
	"documentService/utils"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
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
func (ds *DocumentService) DownloadAllAuthorized(authUser *jwt.Token) (*[]byte, *customerror.CustomError) {
	role := utils.GetUserRole(authUser)
	paths := new([]model.DownloadModel)
	var err *customerror.CustomError
	if role == "User" {
		userId := utils.GetUserId(authUser)
		paths, err = ds.MongoService.GetAllPaths(&userId)
	} else {
		paths, err = ds.MongoService.GetAllPaths(nil)
	}
	if err != nil {
		return nil, err
	}

	zipPath, err := utils.CreateZip(paths)
	if err != nil {
		return nil, err
	}
	fileBytes, errB := ioutil.ReadFile(*zipPath)
	if errB != nil {
		return nil, customerror.NewError(errB.Error(), 500)
	}
	if err := os.RemoveAll(*zipPath); err != nil {
		return nil, customerror.NewError(err.Error(), 500)
	}
	return &fileBytes, nil

}
func (ds *DocumentService) DownloadWithId(id string, authUser *jwt.Token) (*string, *customerror.CustomError) {
	userRole := utils.GetUserRole(authUser)
	path := new(string)
	err := new(customerror.CustomError)
	if userRole == "User" {
		userId := utils.GetUserId(authUser)
		path, err = ds.MongoService.GetPathById(&id, &userId)
	} else {
		path, err = ds.MongoService.GetPathById(&id, nil)
	}
	if err != nil {
		return nil, err
	}
	return path, nil
}
