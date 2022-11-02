package usecase

import (
	"jwt_clean/domain"

	"github.com/gin-gonic/gin"
)

type productUseCase struct {
	productRepo domain.ProductRepository
}

func NewProductUseCase(repo domain.ProductRepository) domain.ProductUseCase {
	return productUseCase{
		productRepo: repo,
	}
}

func (c productUseCase) CreateProductUc(ctx *gin.Context) (*domain.Product, error) {
	return c.productRepo.CreateProductRepository(ctx)
}
