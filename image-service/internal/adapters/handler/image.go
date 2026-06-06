package handler

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".webp": true,
}

func UploadImage(c *gin.Context) {

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado: falta el token de autenticación"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se recibió ninguna imagen"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato no permitido. Use: jpg, jpeg, png, gif, webp"})
		return
	}

	const maxSize = 10 << 20 //30mb maximoas
	if file.Size > maxSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La imagen supera el límite de 10 MB"})
		return
	}

	filename := filepath.Base(file.Filename)
	destination := "uploads/" + filename

	if err := c.SaveUploadedFile(file, destination); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fallo interno al guardar la imagen"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Imagen guardada exitosamente",
		"filename": filename,
	})
}
