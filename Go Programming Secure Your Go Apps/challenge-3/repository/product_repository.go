package repository

import "challenge-3/models"

type ProductRepository interface {
	FindById(id string) *models.Product
}
