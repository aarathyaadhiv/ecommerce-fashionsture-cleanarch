package interfaces

import (
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
)

type ProductRepository interface {
	AddProduct(product models.AddProduct,sellingPrice float64) error
	UpdateProduct(product models.ProductUpdate,id uint) error
	DeleteProduct(id uint) error
	AddCategory(category models.AddCategory) error
	DeleteCategory(id uint) error
	ShowAll(page, count int) ([]models.ProductResponse, error)
	ShowProduct(id uint) (models.ProductResponse, error)
	AddBrand(brand models.AddBrand) error
	DeleteBrand(id uint) error
	FetchProductDetails(productId uint)( models.Product,error)
	ShowCategory(page,count int)([]domain.Category,error)
	ShowBrand(page,count int)([]domain.Brand,error)
	ProductByCategory(id uint,page,count int)([]models.ProductResponse,error)
	ProductByBrand(id uint,page,count int)([]models.ProductResponse,error)
	ProductSearch(word string,page,count int)([]models.ProductResponse,error)
	Quantity(id uint)(uint,error)
	UpdateCategory(update models.UpdateCategory,id uint)error
	UpdateBrand(update models.UpdateBrand,id uint)error
	FetchProductDetailsToAdmin(page,count int)([]models.ProductResponseToAdmin,error)
}
