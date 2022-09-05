package controller

import (
	"documentService/handler"
	"documentService/utils"
	"github.com/golang-jwt/jwt"
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
	authUser := c.Get("user").(*jwt.Token)
	if err := utils.Authorize(authUser, "Admin", "User"); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(500, err.Error())
	}
	result, cErr := dc.DocumentService.UploadMultipleFiles(form, authUser)
	if err != nil {
		return c.JSON(cErr.StatusCode, cErr)
	}
	return c.JSON(201, result)
}
func (dc *DocumentController) ShowAll(c echo.Context) error {
	authUser := c.Get("user").(*jwt.Token)
	if err := utils.Authorize(authUser, "Admin", "User", "Viewer"); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	entities, err := dc.DocumentService.ShowAllEntity(authUser)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, entities)
}
func (dc *DocumentController) Delete(c echo.Context) error {
	authUser := c.Get("user").(*jwt.Token)
	if err := utils.Authorize(authUser, "Admin", "User"); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	id := c.Param("id")
	if err := dc.DocumentService.DeleteEntity(id, authUser); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(200, nil)

}
