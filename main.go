package main

import (
	"assignment-final/controller"
	"assignment-final/lib"
	"assignment-final/model/domain"
	"assignment-final/router"
	"assignment-final/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine := gin.Default()
	db, err := lib.InitDatabase()
	if err != nil {
		fmt.Println("DB Connection Error")
	}
	if err := db.AutoMigrate(&domain.User{}, &domain.Photo{}, &domain.Comment{}, &domain.SocialMedia{}); err != nil {
		log.Fatal(err.Error())
	}

	router.NewUserRouter(*controller.NewAuthController(*service.NewUserService(db))).AuthRouter(ginEngine)

	ginEngine.Run("localhost:8000")
}
