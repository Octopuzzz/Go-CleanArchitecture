package main

import (
	"fmt"
	"jwt_clean/domain"
	producthandler "jwt_clean/product/delivery/http"
	productrepository "jwt_clean/product/repository/gorm"
	productusecase "jwt_clean/product/usecase"
	userhandler "jwt_clean/user/delivery/http"
	userrepository "jwt_clean/user/repository/gorm"
	userusecase "jwt_clean/user/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	APP_PORT := os.Getenv("APP_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", DB_HOST, DB_USERNAME, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.Debug().AutoMigrate(domain.User{})
	db.Debug().AutoMigrate(domain.Product{})

	router := gin.Default()

	userRepo := userrepository.NewUserRepository(db)
	userUseCase := userusecase.NewUserUseCase(userRepo)
	userhandler.NewUserHandler(router, userUseCase)

	productRepo := productrepository.NewProductRepository(db)
	productUseCase := productusecase.NewProductUseCase(productRepo)
	producthandler.NewProductHandler(router, productUseCase)

	router.Run(fmt.Sprintf(":%s", APP_PORT))
	log.Println("Berjalan pada port :", APP_PORT)
}
