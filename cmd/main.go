package main

import (
	"documentService/config"
	"documentService/controller"
	"documentService/handler"
	"documentService/repository/mongodb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.GET("/api/dms/show", documentController.ShowAll, middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("goapisecretkey")}))
	e.POST("/api/dms/upload", documentController.Upload, middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("goapisecretkey")}))
	e.DELETE("/api/dms/delete/:id", documentController.Delete, middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("goapisecretkey")}))

	if err := e.Start(":5000"); err != nil {
		panic(err)
	}
}
