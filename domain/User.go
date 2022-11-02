package domain

import (
	"jwt_clean/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string    `json:"full_name" gorm:"not null" form:"full_name" valid:"required~Your full name is required"`
	Email    string    `json:"email" gorm:"not null;uniqueIndex" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string    `json:"password" gorm:"not null" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 5 characters"`
	Product  []Product `json:"products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(user)

	if errCreate != nil {
		err = errCreate
		return
	}

	user.Password = helpers.HashPass(user.Password)
	err = nil
	return
}

type UserUseCase interface {
	GetUserByIdUc(ctx *gin.Context) (User, error)
	UserRegisterUc(ctx *gin.Context) (*User, error)
	UserLoginUc(ctx *gin.Context) error
}

type UserRepository interface {
	GetUserByIdRepository(ctx *gin.Context) (User, error)
	UserRegisterRepository(ctx *gin.Context) (*User, error)
	UserLoginRepository(ctx *gin.Context) error
}
