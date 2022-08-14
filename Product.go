package Products_Age

import (
	"time"
)

type Product struct {
	name         string
	category     string
	weight       float32
	price        float64
	creationDate time.Time
}

func (product Product) GetCreationDate() time.Time {
	return product.creationDate
}
