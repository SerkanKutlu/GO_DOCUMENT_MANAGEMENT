package handler

import "documentService/repository/mongodb"

type DocumentService struct {
	MongoService *mongodb.MongoService
}

func GetDocumentService(mongoService *mongodb.MongoService) *DocumentService {
	return &DocumentService{MongoService: mongoService}
}
