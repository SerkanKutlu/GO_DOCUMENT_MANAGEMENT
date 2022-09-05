package model

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
	"os"
	"time"
)

type Document struct {
	Id               string `bson:"_id" json:"_id"`
	UserId           string
	UploadedAt       time.Time
	ExpireAt         time.Time
	OriginalFileName string
	FileName         string
	MimeType         string
	Path             string
}

func CreateDocumentEntity(file *multipart.FileHeader, userId string) *Document {
	wd, _ := os.Getwd()
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	path := fmt.Sprintf("%s\\documents\\%s", wd, fileName)
	return &Document{
		Id:               uuid.NewV4().String(),
		UserId:           userId,
		UploadedAt:       time.Now(),
		ExpireAt:         time.Now().Add(30 * 24 * time.Hour),
		OriginalFileName: file.Filename,
		FileName:         fileName,
		MimeType:         file.Header.Get("Content-Type"),
		Path:             path,
	}
}
