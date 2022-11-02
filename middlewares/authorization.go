package middlewares

// import (
// 	"go-jwt/database"
// 	"go-jwt/models"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v4"
// )

// func ProductAuthorization() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		db := database.GetDB()
// 		productId, err := strconv.Atoi(c.Param("productId"))
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 				"error":   "Bad Request",
// 				"message": "invalid parameter",
// 			})
// 			return
// 		}
// 		userData := c.MustGet("userData").(jwt.MapClaims)
// 		userID := uint(userData["id"].(float64))
// 		Product := models.Product{}

// 		err = db.Select("user_id").First(&Product, uint(productId)).Error

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 				"error":   "Data not found",
// 				"message": "data doesn't exist",
// 			})
// 			return
// 		}

// 		if Product.UserID != userID {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error":   "Unauthorized",
// 				"message": "you're not allowed to access this data",
// 			})
// 			return
// 		}
// 		c.Next()
// 	}
// }
