package internal

import (
	"time"
)

type Product struct {
	Name         string
	Category     string
	Weight       float32
	Price        float64
	CreationDate time.Time
}

func (product Product) GetCreationDate() time.Time {
	return product.CreationDate
}
