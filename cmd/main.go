package main

import (
	"documentService/config"
	"documentService/controller"
	"documentService/handler"
	"documentService/repository/mongodb"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	env := os.Getenv("GO_ENV")
	confManager := config.NewConfigurationManager("yml", "application", env)
	mongoConfig := confManager.GetMongoConfiguration()
	mongoService := mongodb.GetMongoService(mongoConfig)
	documentService := handler.GetDocumentService(mongoService)
	documentController := controller.GetDocumentController(documentService)
	e := echo.New()
	e.POST("/api/dms/upload", documentController.Upload)

	if err := e.Start(":5000"); err != nil {
		panic(err)
	}
}
