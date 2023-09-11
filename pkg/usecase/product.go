package usecase

import (
	"fmt"
	"strconv"

	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/helper"
	repo "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	services "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
)

type ProductUseCase struct {
	ProductRepo repo.ProductRepository
}

func NewProductUseCase(repo repo.ProductRepository) services.ProductUseCase {
	return &ProductUseCase{repo}
}

func (c *ProductUseCase) AddProduct(product models.AddProduct) error {
	sellingPrice := helper.SellingPrice(product.Price, product.Discount)
	return c.ProductRepo.AddProduct(product, sellingPrice)
}

func (c *ProductUseCase) UpdateProduct(product models.ProductUpdate) error {
	return c.ProductRepo.UpdateProduct(product)
}
func (c *ProductUseCase) DeleteProduct(id string) error {
	productID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return c.ProductRepo.DeleteProduct(uint(productID))
}

func (c *ProductUseCase) AddCategory(category models.AddCategory) error {
	return c.ProductRepo.AddCategory(category)
}

func (c *ProductUseCase) DeleteCategory(id string) error {
	categoryID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return c.ProductRepo.DeleteCategory(uint(categoryID))
}

func (c *ProductUseCase) ShowAll(page, count int) ([]models.ProductResponse, error) {
	productResponse, err := c.ProductRepo.ShowAll(page, count)
	if err != nil {
		return nil, err
	}
	return productResponse, nil

}

func (c *ProductUseCase) ShowProduct(id string) (models.ProductResponse, error) {
	productId, err := strconv.Atoi(id)

	if err != nil {
		return models.ProductResponse{}, err
	}
	productResponse, err := c.ProductRepo.ShowProduct(uint(productId))

	if err != nil {
		return models.ProductResponse{}, nil
	}
	return productResponse, nil

}

func (c *ProductUseCase) AddBrand(brand models.AddBrand) error {
	return c.ProductRepo.AddBrand(brand)
}

func (c *ProductUseCase) DeleteBrand(id string) error {
	brandId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return c.ProductRepo.DeleteBrand(uint(brandId))
}

func (c *ProductUseCase) ShowCategory() ([]domain.Category, error) {
	return c.ProductRepo.ShowCategory()
}

func (c *ProductUseCase) ShowBrand() ([]domain.Brand, error) {
	return c.ProductRepo.ShowBrand()
}

func (c *ProductUseCase) FilterProductByCategory(id string) ([]models.ProductResponse, error) {
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return c.ProductRepo.ProductByCategory(uint(categoryId))
}

func (c *ProductUseCase) FilterProductByBrand(id string) ([]models.ProductResponse, error) {
	brandId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return c.ProductRepo.ProductByBrand(uint(brandId))
}

func (c *ProductUseCase) ProductSearch(word string)([]models.ProductResponse,error){
	words:=fmt.Sprint(word)
	searchWord:="%"+words+"%"
	return c.ProductRepo.ProductSearch(searchWord)
}