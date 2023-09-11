package interfaces

import (
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
)

type ProductUseCase interface {
	AddProduct(product models.AddProduct) error
	UpdateProduct(product models.ProductUpdate) error
	DeleteProduct(id string) error
	AddCategory(category models.AddCategory) error
	DeleteCategory(id string) error
	ShowAll(page, count int) ([]models.ProductResponse, error)
	ShowProduct(id string) (models.ProductResponse, error)
	AddBrand(brand models.AddBrand) error
	DeleteBrand(id string) error
	ShowCategory()([]domain.Category,error)
	ShowBrand()([]domain.Brand,error)
	FilterProductByCategory(id string)([]models.ProductResponse,error)
	FilterProductByBrand(id string)([]models.ProductResponse,error)
	ProductSearch(word string)([]models.ProductResponse,error)
}
