package controller

import (
	"documentService/handler"
	"github.com/labstack/echo/v4"
)

type DocumentController struct {
	DocumentService *handler.DocumentService
}

func GetDocumentController(documentService *handler.DocumentService) *DocumentController {
	return &DocumentController{
		documentService,
	}
}

func (dc *DocumentController) Upload(c echo.Context) error {

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(500, err.Error())
	}
	xx, _ := dc.DocumentService.UploadMultipleFiles(form)
	return c.JSON(200, xx)
}
