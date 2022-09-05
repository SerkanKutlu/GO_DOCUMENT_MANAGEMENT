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
	result, err := dc.DocumentService.UploadMultipleFiles(form)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, result)
}
