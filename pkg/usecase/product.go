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

func (c *ProductUseCase) UpdateProduct(product models.ProductUpdate,id string) error {
	productId,err:=strconv.Atoi(id)
	if err!=nil{
		return err
	}
	return c.ProductRepo.UpdateProduct(product,uint(productId))
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
	updatedProductResponse:=make([]models.ProductResponse,0)
	
	for _,product:=range productResponse{
		quantity,_:=c.ProductRepo.Quantity(product.ID)
		if quantity==0{
			product.Status="out of stock"
		}else if quantity==1{
			product.Status="only 1 product remains"
		}else{
			product.Status="in stock"
		}
		updatedProductResponse = append(updatedProductResponse, product)
	}
	return updatedProductResponse, nil

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
	quantity,_:=c.ProductRepo.Quantity(productResponse.ID)
	if quantity==0{
		productResponse.Status="out of stock"
	}else if quantity==1{
		productResponse.Status="only 1 product remains"
	}else{
		productResponse.Status="in stock"
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

func (c *ProductUseCase) ShowCategory(pages,counts string) ([]domain.Category, error) {
	page,err:=strconv.Atoi(pages)
	if err!=nil{
		return nil,err
	}
	count,err:=strconv.Atoi(counts)
	if err!=nil{
		return nil,err
	}
	return c.ProductRepo.ShowCategory(page,count)
}

func (c *ProductUseCase) ShowBrand(pages,counts string) ([]domain.Brand, error) {
	page,err:=strconv.Atoi(pages)
	if err!=nil{
		return nil,err
	}
	count,err:=strconv.Atoi(counts)
	if err!=nil{
		return nil,err
	}
	return c.ProductRepo.ShowBrand(page,count)
}

func (c *ProductUseCase) FilterProductByCategory(id string,pages,counts string) ([]models.ProductResponse, error) {
	page,err:=strconv.Atoi(pages)
	if err!=nil{
		return nil,err
	}
	count,err:=strconv.Atoi(counts)
	if err!=nil{
		return nil,err
	}
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return c.ProductRepo.ProductByCategory(uint(categoryId),page,count)
}

func (c *ProductUseCase) FilterProductByBrand(id string,pages,counts string) ([]models.ProductResponse, error) {
	page,err:=strconv.Atoi(pages)
	if err!=nil{
		return nil,err
	}
	count,err:=strconv.Atoi(counts)
	if err!=nil{
		return nil,err
	}
	brandId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return c.ProductRepo.ProductByBrand(uint(brandId),page,count)
}

func (c *ProductUseCase) ProductSearch(word string,pages,counts string)([]models.ProductResponse,error){
	page,err:=strconv.Atoi(pages)
	if err!=nil{
		return nil,err
	}
	count,err:=strconv.Atoi(counts)
	if err!=nil{
		return nil,err
	}
	words:=fmt.Sprint(word)
	searchWord:="%"+words+"%"
	return c.ProductRepo.ProductSearch(searchWord,page,count)
}

func (c *ProductUseCase) UpdateCategory(update models.UpdateCategory,id string)error{
	categoryId,err:=strconv.Atoi(id)
	if err!=nil{
		return err
	}
	return c.ProductRepo.UpdateCategory(update,uint(categoryId))
}

func (c *ProductUseCase) UpdateBrand(update models.UpdateBrand,id string)error{
	brandId,err:=strconv.Atoi(id)
	if err!=nil{
		return err
	}
	return c.ProductRepo.UpdateBrand(update,uint(brandId))
}

func (c *ProductUseCase) GetProductToAdmin(pages,counts string)([]models.ProductResponseToAdmin,error){
	page,err:=strconv.Atoi(pages)
	if err!=nil{
		return nil,err
	}
	count,err:=strconv.Atoi(counts)
	if err!=nil{
		return nil,err
	}
	return c.ProductRepo.FetchProductDetailsToAdmin(page,count)
}