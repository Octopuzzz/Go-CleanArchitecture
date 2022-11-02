package http

import (
	"jwt_clean/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUseCase domain.ProductUseCase
}

func NewProductHandler(r *gin.Engine, productUc domain.ProductUseCase) {
	handler := &ProductHandler{
		productUseCase: productUc,
	}
	router := r.Group("/users")
	router.POST("/", handler.CreateProductUc)
}

func (h ProductHandler) CreateProductUc(c *gin.Context) {
	res, err := h.productUseCase.CreateProductUc(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, res)
}
