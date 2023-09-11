package interfaces

import (
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
)

type ProductRepository interface {
	AddProduct(product models.AddProduct,sellingPrice float64) error
	UpdateProduct(product models.ProductUpdate) error
	DeleteProduct(id uint) error
	AddCategory(category models.AddCategory) error
	DeleteCategory(id uint) error
	ShowAll(page, count int) ([]models.ProductResponse, error)
	ShowProduct(id uint) (models.ProductResponse, error)
	AddBrand(brand models.AddBrand) error
	DeleteBrand(id uint) error
	FetchProductDetails(productId uint)( models.Product,error)
	ShowCategory()([]domain.Category,error)
	ShowBrand()([]domain.Brand,error)
	ProductByCategory(id uint)([]models.ProductResponse,error)
	ProductByBrand(id uint)([]models.ProductResponse,error)
	ProductSearch(word string)([]models.ProductResponse,error)
}
