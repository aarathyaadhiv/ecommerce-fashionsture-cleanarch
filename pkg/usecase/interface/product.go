package interfaces

import (
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
)

type ProductUseCase interface {
	AddProduct(product models.AddProduct) error
	UpdateProduct(product models.ProductUpdate,id string) error
	DeleteProduct(id string) error
	AddCategory(category models.AddCategory) error
	DeleteCategory(id string) error
	ShowAll(page, count int) ([]models.ProductResponse, error)
	ShowProduct(id string) (models.ProductResponse, error)
	AddBrand(brand models.AddBrand) error
	DeleteBrand(id string) error
	ShowCategory(pages,counts string)([]domain.Category,error)
	ShowBrand(pages,counts string)([]domain.Brand,error)
	FilterProductByCategory(id string,pages,counts string)([]models.ProductResponse,error)
	FilterProductByBrand(id string,pages,counts string)([]models.ProductResponse,error)
	ProductSearch(word string,pages,counts string)([]models.ProductResponse,error)
	UpdateCategory(update models.UpdateCategory,id string)error
	UpdateBrand(update models.UpdateBrand,id string)error
	GetProductToAdmin(pages,counts string)([]models.ProductResponseToAdmin,error)
}
