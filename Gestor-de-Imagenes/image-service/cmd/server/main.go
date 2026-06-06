package main

import (
	"fmt"
	"image-service/internal/adapters/handler"
	"image-service/internal/core/domain"
	"image-service/internal/middleware"
	"io"
	"net/http"
	"time"
)

type mockService struct{}

func (s *mockService) UploadImage(name string, size int64, file io.Reader, userID string) (*domain.Image, error) {
	return &domain.Image{
		ID:         "img_98765",
		Name:       name,
		Size:       size,
		Format:     "image/png",
		UserID:     userID,
		UploadedAt: time.Now(),
	}, nil
}

func main() {
	service := &mockService{}
	httpHandler := handler.NewHTTPHandler(service)

	protectedHandler := middleware.RequireAuth(httpHandler.UploadHandler)

	http.HandleFunc("/api/v1/images/upload", protectedHandler)

	port := ":5000"
	fmt.Printf("Servidor de Gestión de Imágenes (Puerto %s)...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
