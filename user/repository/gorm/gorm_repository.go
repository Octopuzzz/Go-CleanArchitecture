package gorm

import (
	"jwt_clean/domain"
	"jwt_clean/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (m *userRepository) GetUserByIdRepository(c *gin.Context) (user domain.User, err error) {
	id := c.Param("userId")
	err = m.DB.Model(&user).Where("id=?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (m *userRepository) UserRegisterRepository(c *gin.Context) (user *domain.User, err error) {
	contentType := helpers.GetContentType(c)
	if contentType == "application/json" {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}
	err = m.DB.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *userRepository) UserLoginRepository(c *gin.Context) (err error) {
	var user domain.User
	contentType := helpers.GetContentType(c)

	if contentType == "application/json" {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	password := user.Password
	err = m.DB.Debug().Where("email=?", user.Email).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "Invalid email/password",
		})
		return err
	}
	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "Invalid email/password",
		})
		return err
	}
	token := helpers.GenerateToken(user.ID, user.Email)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
	return nil
}
