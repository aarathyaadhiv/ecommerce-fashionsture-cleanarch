package repository

import (
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	repo "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) repo.ProductRepository {
	return &ProductRepository{DB}
}

func (c *ProductRepository) AddProduct(product models.AddProduct,sellingPrice float64) error {
	return c.DB.Exec(`insert into products (name,description,quantity,price,selling_price,image,discount,category_id,brand_id) values(?,?,?,?,?,?,?,?,?)`, product.Name, product.Description, product.Quantity, product.Price,sellingPrice, product.Image, product.Discount, product.CategoryID, product.BrandID).Error
}

func (c *ProductRepository) UpdateProduct(product models.ProductUpdate) error {
	return c.DB.Exec(`update products set quantity=? where id=?`, product.Quantity, product.ID).Error
}

func (c *ProductRepository) DeleteProduct(id uint) error {
	return c.DB.Exec(`delete from products where id=?`, id).Error
}

func (c *ProductRepository) AddCategory(category models.AddCategory) error {
	return c.DB.Exec(`insert into categories(name) values(?)`, category.Name).Error
}

func (c *ProductRepository) DeleteCategory(id uint) error {
	return c.DB.Exec(`delete from categories where id=?`, id).Error
}

func (c *ProductRepository) ShowAll(page, count int) ([]models.ProductResponse, error) {
	offset := (page - 1) * count

	var productResponse []models.ProductResponse
	if err := c.DB.Raw(`select p.id,p.name,p.description,p.price,p.selling_price,p.image,p.discount,c.name as category,b.name as brand from products p join categories c on c.id=p.category_id join brands b on b.id=p.brand_id limit ? offset ?`, count, offset).Scan(&productResponse).Error; err != nil {
		return nil, err
	}
	return productResponse, nil
}

func (c *ProductRepository) ShowProduct(id uint) (models.ProductResponse, error) {
	var productDetails models.ProductResponse

	if err := c.DB.Raw(`select p.id,p.name,p.description,p.price,p.selling_price,p.image,p.discount,c.name as category,b.name as brand from products p join categories c on c.id=p.category_id join brands b on b.id=p.brand_id where p.id=?`, id).Scan(&productDetails).Error; err != nil {
		return models.ProductResponse{}, err
	}
	return productDetails, nil
}

func (c *ProductRepository) AddBrand(brand models.AddBrand) error {
	return c.DB.Exec(`insert into brands(name) values(?)`, brand.Name).Error
}

func (c *ProductRepository) DeleteBrand(id uint) error {
	return c.DB.Exec(`delete from brands where id=?`, id).Error
}

func (c *ProductRepository) FetchProductDetails(productId uint)( models.Product,error){
	var product models.Product
	err:=c.DB.Raw(`SELECT selling_price,quantity FROM products WHERE id=?`,productId).Scan(&product).Error
	return product,err
}

func (c *ProductRepository) ShowCategory()([]domain.Category,error){
	var category []domain.Category
	err:=c.DB.Raw(`SELECT * FROM categories`).Scan(&category).Error
	if err!=nil{
		return nil,err
	}
	return category,nil
}

func (c *ProductRepository) ShowBrand()([]domain.Brand,error){
	var brand []domain.Brand
	err:=c.DB.Raw(`SELECT * FROM brands`).Scan(&brand).Error
	if err!=nil{
		return nil,err
	}
	return brand,nil
}

func (c *ProductRepository) ProductByCategory(id uint)([]models.ProductResponse,error){
	var product []models.ProductResponse
	err:=c.DB.Raw(`select p.id,p.name,p.description,p.price,p.selling_price,p.image,p.discount,c.name as category,b.name as brand from products p join categories c on c.id=p.category_id join brands b on b.id=p.brand_id where p.category_id=?`,id).Scan(&product).Error
	if err!=nil{
		return nil,err
	}
	return product,nil
}

func (c *ProductRepository) ProductByBrand(id uint)([]models.ProductResponse,error){
	var product []models.ProductResponse
	err:=c.DB.Raw(`select p.id,p.name,p.description,p.price,p.selling_price,p.image,p.discount,c.name as category,b.name as brand from products p join categories c on c.id=p.category_id join brands b on b.id=p.brand_id where p.brand_id=?`,id).Scan(&product).Error
	if err!=nil{
		return nil,err
	}
	return product,nil
}

func (c *ProductRepository) ProductSearch(word string)([]models.ProductResponse,error){
	var product []models.ProductResponse
	err:=c.DB.Raw(`select p.id,p.name,p.description,p.price,p.selling_price,p.image,p.discount,c.name as category,b.name as brand from products p join categories c on c.id=p.category_id join brands b on b.id=p.brand_id where p.name ILIKE ?`,word).Scan(&product).Error
	if err!=nil{
		return nil,err
	}
	return product,nil
}

func (c *ProductRepository) Quantity(id uint)(uint,error){
	var quantity uint
	err:=c.DB.Raw(`SELECT quantity FROM products WHERE id=?`,id).Scan(&quantity).Error
	if err!=nil{
		return 0,err
	}
	return quantity,err
}