package http

import (
	"jwt_clean/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase domain.UserUseCase
}

func NewUserHandler(r *gin.Engine, userUc domain.UserUseCase) {
	handler := &UserHandler{
		userUseCase: userUc,
	}
	router := r.Group("/users")
	router.POST("/register", handler.UserRegister)
	router.POST("/login", handler.UserLogin)
	router.GET("/:userId", handler.GetUserById)
}
func (h UserHandler) GetUserById(c *gin.Context) {
	var result gin.H
	res, err := h.userUseCase.GetUserByIdUc(c)
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
		c.JSON(http.StatusNotFound, result)
		return
	} else {
		result = gin.H{
			"result": res,
			"count":  1,
		}
	}
	c.JSON(http.StatusNotFound, result)
}

func (h UserHandler) UserRegister(c *gin.Context) {
	res, err := h.userUseCase.UserRegisterUc(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":        res.ID,
		"email":     res.Email,
		"full_name": res.FullName,
	})
}

func (h UserHandler) UserLogin(c *gin.Context) {
	err := h.userUseCase.UserLoginUc(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusCreated, err)
}
