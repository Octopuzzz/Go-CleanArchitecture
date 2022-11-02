package gorm

import (
	"jwt_clean/domain"
	"jwt_clean/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{
		DB: db,
	}
}

func (m *productRepository) CreateProductRepository(c *gin.Context) (product *domain.Product, err error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	userID := uint(userData["id"].(float64))
	if contentType == "application/json" {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}
	product.UserID = userID
	err = m.DB.Debug().Create(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (m *productRepository) BeforeCreate(tx *gorm.DB) (product domain.Product, err error) {
	_, errCreate := govalidator.ValidateStruct(product)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
func (m *productRepository) BeforeUpdate(tx *gorm.DB) (product domain.Product, err error) {
	_, errCreate := govalidator.ValidateStruct(product)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
