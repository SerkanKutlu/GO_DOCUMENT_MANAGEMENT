package controller

import (
	"documentService/handler"
	"documentService/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
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
func (dc *DocumentController) ShowAllAuthorized(c echo.Context) error {
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
func (dc *DocumentController) DownloadAllAuthorized(c echo.Context) error {
	utils.CreateZip()
	fileBytes, err := ioutil.ReadFile("downloading.zip")
	if err != nil {
		panic(err)
	}
	c.Response().Header().Set("Content-Type", "application/zip")
	c.Response().WriteHeader(200)
	c.Response().Writer.Write(fileBytes)
	return nil
	//authUser := c.Get("user").(*jwt.Token)
	//if err := utils.Authorize(authUser, "Admin", "User", "Viewer"); err != nil {
	//	return c.JSON(err.StatusCode, err)
	//}
	//file, err := dc.DocumentService.DownloadAllAuthorized(authUser)
	//if err != nil {
	//	return c.JSON(err.StatusCode, err)
	//}
	//list := new([]interface{})
	//*list = append(*list, c.File(*file))
	//*list = append(*list, c.File(*file))
	//*list = append(*list, c.File(*file))
	//*list = append(*list, c.File(*file))
	//
	//return c.JSON(200, list)

}
func (dc *DocumentController) DownloadWithId(c echo.Context) error {
	authUser := c.Get("user").(*jwt.Token)
	if err := utils.Authorize(authUser, "Admin", "User", "Viewer"); err != nil {
		return c.JSON(err.StatusCode, err)
	}
	id := c.Param("id")
	path, err := dc.DocumentService.DownloadWithId(id, authUser)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	errb := c.File(*path)
	if errb != nil {
		return c.JSON(400, errb.Error())
	}
	return c.JSON(200, c.File(*path))
}
