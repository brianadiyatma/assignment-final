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

	///Dependency Injection User & Auth
	userService := *service.NewUserService(db)
	authController :=*controller.NewAuthController(&userService)
	userController :=*controller.NewUserController(&userService)
	router.NewUserRouter(authController, userController).AuthRouter(ginEngine)
	//Dependency Injection Photo
	photoService := *service.NewPhotoService(db)
	photoController := *controller.NewPhotoController(&photoService)
	router.NewPhotoRouter(photoController).PhotoRouter(ginEngine)
	//Dependency Injection Sosmed
	socialMediaService := *service.NewSocialMedia(db)
	socialMediaController:=*controller.NewSocialMediaController(&socialMediaService)
	router.NewSocialMediaRouter(socialMediaController).SocialMediaRouter(ginEngine)
	//Dependency Injection Comment
	commentService := *service.NewCommentService(db)
	commentController := *controller.NewCommentService(&commentService)
	router.NewCommentRouter(commentController).CommentRouter(ginEngine)



	ginEngine.Run("localhost:8000")
}
