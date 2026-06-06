package ports

import (
	"image-service/internal/core/domain"
	"io"
)

type ImageService interface {
	UploadImage(name string, size int64, file io.Reader, userID string) (*domain.Image, error)
}

type ImageRepository interface {
	Save(image *domain.Image, file io.Reader) error
}
