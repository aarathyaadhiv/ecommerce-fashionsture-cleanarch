package handler

import (
	"net/http"
	"strconv"

	handler "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/handler/interface"
	services "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Usecase services.ProductUseCase
}

func NewProductHandler(usecase services.ProductUseCase) handler.ProductHandler {
	return &ProductHandler{usecase}
}

// @Summary Add Product
// @Description Add Product By Admin
// @Tags Product Management
// @Accept json
// @Produce json
// @Param  product body models.AddProduct true "product details"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/product [post]
func (pr *ProductHandler) AddProduct(c *gin.Context) {
	var addProduct models.AddProduct

	if err := c.ShouldBindJSON(&addProduct); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are not in the required format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err := pr.Usecase.AddProduct(addProduct)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully added product", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Update Product
// @Description Update Product By Admin
// @Tags Product Management
// @Accept json
// @Produce json
// @Param  product body models.ProductUpdate true "update details"
// @Param  id path string true "id"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/product/{id} [patch]
func (pr *ProductHandler) UpdateProduct(c *gin.Context) {
	var updateProduct models.ProductUpdate

	if err := c.ShouldBindJSON(&updateProduct); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are not in correct manner", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	id:=c.Param("id")
	err := pr.Usecase.UpdateProduct(updateProduct,id)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully update product", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Delete Product
// @Description Delete Product By Admin
// @Tags Product Management
// @Accept json
// @Produce json
// @Param  id path string true "id"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/product/{id} [delete]
func (pr *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := pr.Usecase.DeleteProduct(id); err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully deleted", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Add Category
// @Description Add Category By Admin
// @Tags Category Management
// @Accept json
// @Produce json
// @Param  category body models.AddCategory true "category details"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/category [post]
func (pr *ProductHandler) AddCategory(c *gin.Context) {
	var addCategory models.AddCategory

	if err := c.ShouldBindJSON(&addCategory); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are not provided in required order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := pr.Usecase.AddCategory(addCategory)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully added category", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Delete Category
// @Description Delete Category By Admin
// @Tags Category Management
// @Accept json
// @Produce json
// @Param  id path string true "id"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/category/{id} [delete]
func (pr *ProductHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	err := pr.Usecase.DeleteCategory(id)

	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully deleted category", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Show All
// @Description Show All Product To User
// @Tags Product View
// @Accept json
// @Produce json
// @Param  page query string true "page"
// @Param  count query string true "count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /products [get]
func (pr *ProductHandler) ShowAll(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil {
		errRes := response.Responses(http.StatusBadRequest, "bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	count, err := strconv.Atoi(c.DefaultQuery("count", "4"))
	if err != nil {
		errRes := response.Responses(http.StatusBadRequest, "bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	productDetails, err := pr.Usecase.ShowAll(page, count)

	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing products", productDetails, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Show Product
// @Description Show Product To User
// @Tags Product View
// @Accept json
// @Produce json
// @Param  id path string true "id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /products/{id} [get]
func (pr *ProductHandler) ShowProduct(c *gin.Context) {
	id := c.Param("id")

	productDetails, err := pr.Usecase.ShowProduct(id)

	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing the product", productDetails, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Add Brand
// @Description Add Brand By Admin
// @Tags Brand Management
// @Accept json
// @Produce json
// @Param  brand body models.AddBrand true "brand details"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/brand [post]
func (pr *ProductHandler) AddBrand(c *gin.Context) {
	var addBrand models.AddBrand

	if err := c.ShouldBindJSON(&addBrand); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are not in proper format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err := pr.Usecase.AddBrand(addBrand)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully added the brand", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Delete Brand
// @Description Delete Brand By Admin
// @Tags Brand Management
// @Accept json
// @Produce json
// @Param  id path string true "id"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/brand/{id} [delete]
func (pr *ProductHandler) DeleteBrand(c *gin.Context) {
	id := c.Param("id")
	err := pr.Usecase.DeleteBrand(id)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully deleted brand", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Show All categories
// @Description Show All categories To User
// @Tags Filter
// @Accept json
// @Produce json
// @Param  page query string true "page"
// @Param  count query string true "count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /category [get]
func (pr *ProductHandler) ShowCategory(c *gin.Context) {
	page:=c.DefaultQuery("page","1")
	count:=c.DefaultQuery("count","3")
	category, err := pr.Usecase.ShowCategory(page,count)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing category", category, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Show All brands
// @Description Show All brands To User
// @Tags Filter
// @Accept json
// @Produce json
// @Param  page query string true "page"
// @Param  count query string true "count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /brand [get]
func (pr *ProductHandler) ShowBrand(c *gin.Context) {
	page:=c.DefaultQuery("page","1")
	count:=c.DefaultQuery("count","3")
	brand, err := pr.Usecase.ShowBrand(page,count)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing brand", brand, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Filter Products By category
// @Description Filter Products By category
// @Tags Filter
// @Accept json
// @Produce json
// @Param  id path string true "id"
// @Param  page query string true "page"
// @Param  count query string true "count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /products/category/{id} [get]
func (pr *ProductHandler) FilterProductsByCategory(c *gin.Context) {
	id := c.Param("id")
	page:=c.DefaultQuery("page","1")
	count:=c.DefaultQuery("count","3")
	products, err := pr.Usecase.FilterProductByCategory(id,page,count)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing products with given category", products, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Filter Products By Brand
// @Description Filter Products By Brand
// @Tags Filter
// @Accept json
// @Produce json
// @Param  id path string true "id"
// @Param  page query string true "page"
// @Param  count query string true "count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /products/brand/{id} [get]
func (pr *ProductHandler) FilterProductsByBrand(c *gin.Context) {
	id := c.Param("id")
	page:=c.DefaultQuery("page","1")
	count:=c.DefaultQuery("count","3")
	products, err := pr.Usecase.FilterProductByBrand(id,page,count)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing products with given brand", products, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Show Products By name
// @Description Show Products By A Word In The Name Of The Product
// @Tags Product View
// @Accept json
// @Produce json
// @Param  word body models.ProductSearch true "search word"
// @Param  page query string true "page"
// @Param  count query string true "count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /products/search [get]
func (pr *ProductHandler) SearchProduct(c *gin.Context) {
	var word models.ProductSearch
	if err := c.ShouldBindJSON(&word); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are not in proper format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	page:=c.DefaultQuery("page","1")
	count:=c.DefaultQuery("count","3")
	products, err := pr.Usecase.ProductSearch(word.Word,page,count)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing products with given word in name", products, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Update category
// @Description Update categry By Admin
// @Tags Category Management
// @Accept json
// @Produce json
// @Param  updateCategory body models.UpdateCategory true "update category"
// @Param  id path string true "id"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/category/{id} [patch]
func (pr *ProductHandler) UpdateCategory(c *gin.Context) {
	var updateCategory models.UpdateCategory

	if err := c.ShouldBindJSON(&updateCategory); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are not in correct manner", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	id := c.Param("id")
	err := pr.Usecase.UpdateCategory(updateCategory, id)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully updated category", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Update brand
// @Description Update brand By Admin
// @Tags Brand Management
// @Accept json
// @Produce json
// @Param  updateBrand body models.UpdateBrand true "update brand"
// @Param  id path string true "id"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/brand/{id} [patch]
func (pr *ProductHandler) UpdateBrand(c *gin.Context) {
	var updateBrand models.UpdateBrand

	if err := c.ShouldBindJSON(&updateBrand); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are not in correct manner", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	id := c.Param("id")
	err := pr.Usecase.UpdateBrand(updateBrand, id)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully updated brand", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Get Product
// @Description Get Product To Admin
// @Tags Product Management
// @Accept json
// @Produce json
// @Param  updateBrand body models.UpdateBrand true "update brand"
// @Param  page query string true "page"
// @Param  count query string true "count"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/product [get]
func (pr *ProductHandler) GetProductsToAdmin(c *gin.Context){
	page:=c.DefaultQuery("page","1")
	count:=c.DefaultQuery("count","3")
	product,err:=pr.Usecase.GetProductToAdmin(page,count)
	if err!=nil{
		errRes:=response.Responses(http.StatusInternalServerError,"internal serverr error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.Responses(http.StatusOK,"successfully showing products",product,nil)
	c.JSON(http.StatusOK,succRes)
}
