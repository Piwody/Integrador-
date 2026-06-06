package main

import (
	"fmt"
	"image-service/internal/adapters/handler"
	"image-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/api/v1/images/upload", middleware.RequireAuth(), handler.UploadImage)

	port := ":5000"
	fmt.Printf("Servidor de Gestión de Imágenes escuchando en puerto %s...\n", port)
	router.Run(port)
}
