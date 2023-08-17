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
// @Router /admin/product/add [post]
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
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/product/update [patch]
func (pr *ProductHandler) UpdateProduct(c *gin.Context) {
	var updateProduct models.ProductUpdate

	if err := c.ShouldBindJSON(&updateProduct); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are not in correct manner", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := pr.Usecase.UpdateProduct(updateProduct)
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
// @Router /admin/product/delete/{id} [delete]
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
// @Router /admin/category/add [post]
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
// @Router /admin/category/delete/{id} [delete]
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
// @Router /admin/brand/add [post]
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
// @Router /admin/brand/delete/{id} [delete]
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
