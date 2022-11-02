package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string `json:"title" form:"title" valid:"required~Title of your product is reuired"`
	Description string `json:"description" form:"description" valid:"required~Description of your product is required"`
	UserID      uint
	User        *User
}

type ProductUseCase interface {
	CreateProductUc(c *gin.Context) (*Product, error)
}

type ProductRepository interface {
	CreateProductRepository(c *gin.Context) (*Product, error)
}
