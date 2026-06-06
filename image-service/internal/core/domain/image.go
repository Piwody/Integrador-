package domain

import "time"

type Image struct {
	ID         string    `json:"id"`
	Nombre     string    `json:"nombre"`
	Ruta       string    `json:"ruta"`
	Tags       []string  `json:"tags"`
	Resolucion string    `json:"resolucion"`
	Tamano     int64     `json:"tamano"`
	Formato    string    `json:"formato"`
	CreatedAt  time.Time `json:"created_at"`
}
