package main

import (
	"documentService/config"
	"documentService/controller"
	"documentService/customMiddleware"
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

	httpConfig := confManager.GetHttpClientConfig()
	customMiddleware.SetHttpClient(httpConfig)
	jwtKey := confManager.GetJwtKey().SecretKey
	e := echo.New()
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte(jwtKey)}))
	e.Use(customMiddleware.ValidateToken)
	e.GET("/api/dms/show", documentController.ShowAllAuthorized, customMiddleware.UseAuthorization("Admin,User,Viewer"))
	e.POST("/api/dms/upload", documentController.Upload, customMiddleware.UseAuthorization("Admin", "User"))
	e.DELETE("/api/dms/delete/:id", documentController.Delete, customMiddleware.UseAuthorization("Admin", "User"))
	e.GET("/api/dms/download/all", documentController.DownloadAllAuthorized, customMiddleware.UseAuthorization("Admin", "User", "Viewer"))
	e.GET("/api/dms/download/:id", documentController.DownloadWithId, customMiddleware.UseAuthorization("Admin", "User", "Viewer"))

	if err := e.Start(":5000"); err != nil {
		panic(err)
	}
}
