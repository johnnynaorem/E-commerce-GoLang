package http

import (
	"e-commerce/internal/domain/models"
	"e-commerce/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: service}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var newProduct models.Product
	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error in data bind " + err.Error(),
		})
		return
	}
	createdProduct, err := h.ProductService.CreateProduct(&newProduct)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, createdProduct)
}

func (h *ProductHandler) GetAllProduct(ctx *gin.Context) {
	products, err := h.ProductService.FindAllProduct()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductById(ctx *gin.Context) {
	productId := ctx.Query("id")
	product, err := h.ProductService.FindProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, product)
}
