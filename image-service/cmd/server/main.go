package main

import (
	"fmt"
	"image-service/internal/adapters/handler"
	"image-service/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))

	router.POST("/api/v1/images/upload", middleware.RequireAuth(), handler.UploadImage)

	port := ":5000"
	fmt.Printf("Servidor de Gestión de Imágenes escuchando en puerto %s...\n", port)
	router.Run(port)
}
