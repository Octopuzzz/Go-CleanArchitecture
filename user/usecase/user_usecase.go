package usecase

import (
	"jwt_clean/domain"

	"github.com/gin-gonic/gin"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) domain.UserUseCase {
	return userUseCase{
		userRepo: repo,
	}
}

func (c userUseCase) GetUserByIdUc(ctx *gin.Context) (domain.User, error) {
	return c.userRepo.GetUserByIdRepository(ctx)
}

func (c userUseCase) UserRegisterUc(ctx *gin.Context) (*domain.User, error) {
	return c.userRepo.UserRegisterRepository(ctx)
}
func (c userUseCase) UserLoginUc(ctx *gin.Context) error {
	return c.userRepo.UserLoginRepository(ctx)
}
