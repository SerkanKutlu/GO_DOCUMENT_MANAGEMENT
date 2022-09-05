package model

import (
	"fmt"
	"mime/multipart"
	"os"
	"time"
)

type Document struct {
	UserId           string
	UploadedAt       time.Time
	ExpireAt         time.Time
	OriginalFileName string
	FileName         string
	MimeType         string
	Path             string
}

func CreateDocumentEntity(file *multipart.FileHeader) *Document {
	wd, _ := os.Getwd()
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	return &Document{
		UserId:           "TBD",
		UploadedAt:       time.Now(),
		ExpireAt:         time.Now().Add(30 * 24 * time.Hour),
		OriginalFileName: file.Filename,
		FileName:         fileName,
		MimeType:         file.Header.Get("Content-Type"),
		Path:             fmt.Sprintf("%s\\documents", wd),
	}
}
