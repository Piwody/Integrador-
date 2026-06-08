package handler

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"net/http"
	"os"
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

var aesKey = []byte("clave-aes-256bit-proyecto-integr") // 32 bytes exactos

func encryptAESCBC(plaintext string) (string, error) {
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	plaintextBytes := []byte(plaintext)
	padding := blockSize - len(plaintextBytes)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	plaintextBytes = append(plaintextBytes, padtext...)

	ciphertext := make([]byte, blockSize+len(plaintextBytes))
	iv := ciphertext[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[blockSize:], plaintextBytes)

	return hex.EncodeToString(ciphertext), nil
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

	const maxSize = 10 << 20 // 10 MB
	if file.Size > maxSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La imagen supera el límite de 10 MB"})
		return
	}

	// Encriptasisod
	originalName := filepath.Base(file.Filename)
	encryptedName, err := encryptAESCBC(originalName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al encriptar nombre del archivo"})
		return
	}

	secureFilename := encryptedName + ext

	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear directorio de uploads"})
		return
	}

	destination := "uploads/" + secureFilename
	if err := c.SaveUploadedFile(file, destination); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fallo interno al guardar la imagen"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":        "Imagen guardada exitosamente",
		"original_name":  originalName,
		"encrypted_name": secureFilename,
	})
}
