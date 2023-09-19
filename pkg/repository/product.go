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

func (c *ProductRepository) UpdateProduct(product models.ProductUpdate,id uint) error {
	return c.DB.Exec(`update products set quantity=? where id=?`, product.Quantity, id).Error
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

func (c *ProductRepository) ShowCategory(page,count int)([]domain.Category,error){
	offset:=(page-1)*count
	var category []domain.Category
	err:=c.DB.Raw(`SELECT * FROM categories limit ? offset ?`,count,offset).Scan(&category).Error
	if err!=nil{
		return nil,err
	}
	return category,nil
}

func (c *ProductRepository) ShowBrand(page,count int)([]domain.Brand,error){
	offset:=(page-1)*count
	var brand []domain.Brand
	err:=c.DB.Raw(`SELECT * FROM brands limit ? offset ?`,count,offset).Scan(&brand).Error
	if err!=nil{
		return nil,err
	}
	return brand,nil
}

func (c *ProductRepository) ProductByCategory(id uint,page,count int)([]models.ProductResponse,error){
	offset:=(page-1)*count
	var product []models.ProductResponse
	err:=c.DB.Raw(`select p.id,p.name,p.description,p.price,p.selling_price,p.image,p.discount,c.name as category,b.name as brand from products p join categories c on c.id=p.category_id join brands b on b.id=p.brand_id where p.category_id=? limit ? offset ?`,id,count,offset).Scan(&product).Error
	if err!=nil{
		return nil,err
	}
	return product,nil
}

func (c *ProductRepository) ProductByBrand(id uint,page,count int)([]models.ProductResponse,error){
	offset:=(page-1)*count
	var product []models.ProductResponse
	err:=c.DB.Raw(`select p.id,p.name,p.description,p.price,p.selling_price,p.image,p.discount,c.name as category,b.name as brand from products p join categories c on c.id=p.category_id join brands b on b.id=p.brand_id where p.brand_id=? limit ? offset ?`,id,count,offset).Scan(&product).Error
	if err!=nil{
		return nil,err
	}
	return product,nil
}

func (c *ProductRepository) ProductSearch(word string,page,count int)([]models.ProductResponse,error){
	offset:=(page-1)*count
	var product []models.ProductResponse
	err:=c.DB.Raw(`select p.id,p.name,p.description,p.price,p.selling_price,p.image,p.discount,c.name as category,b.name as brand from products p join categories c on c.id=p.category_id join brands b on b.id=p.brand_id where p.name ILIKE ? limit ? offset ?`,word,count,offset).Scan(&product).Error
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

func (c *ProductRepository) UpdateCategory(update models.UpdateCategory,id uint)error{
	return c.DB.Exec(`UPDATE categories SET name=? WHERE id=?`,update.Name,id).Error
}

func (c *ProductRepository) UpdateBrand(update models.UpdateBrand,id uint)error{
	return c.DB.Exec(`UPDATE brands SET name=? WHERE id=?`,update.Name,id).Error
}

func (c *ProductRepository) FetchProductDetailsToAdmin(page,count int)([]models.ProductResponseToAdmin,error){
	var product []models.ProductResponseToAdmin
	offset:=(page-1)*count
	err:= c.DB.Raw(`SELECT id,name,price,selling_price,image,discount,quantity FROM products limit ? offset ?`,count,offset).Scan(&product).Error
	if err!=nil{
		return nil,err
	}
	return product,nil
}

func (c *ProductRepository) IsCategoryExist(name string)(bool,error){
	var count int
	err:=c.DB.Raw(`SELECT COUNT(*) FROM categories WHERE name=?`,name).Scan(&count).Error
	if err!=nil{
		return false,err
	}
	return count>0,nil
}


func (c *ProductRepository) IsBrandExist(name string)(bool,error){
	var count int
	err:=c.DB.Raw(`SELECT COUNT(*) FROM brands WHERE name=?`,name).Scan(&count).Error
	if err!=nil{
		return false,err
	}
	return count>0,nil
}

func (c *ProductRepository) ProductByBrandAndCategory(page,count int,category,brand uint)([]models.ProductResponse,error){
	offset:=(page-1)*count
	var product []models.ProductResponse
	err:=c.DB.Raw(`select p.id,p.name,p.description,p.price,p.selling_price,p.image,p.discount,c.name as category,b.name as brand from products p join categories c on c.id=p.category_id join brands b on b.id=p.brand_id where p.brand_id=? AND p.category_id=? limit ? offset ?`,brand,category,count,offset).Scan(&product).Error
	if err!=nil{
		return nil,err
	}
	return product,nil
}