package router

import (
	"assignment-final/controller"
	"assignment-final/middleware"

	"github.com/gin-gonic/gin"
)

type socialMediaRouter struct {
	socialMediaController controller.SocialMediaController
}

func NewSocialMediaRouter(smc controller.SocialMediaController) *socialMediaRouter {
	return &socialMediaRouter{
		socialMediaController: smc,
	}
}

func (smr *socialMediaRouter) SocialMediaRouter(r *gin.Engine) {
	socialMedias := r.Group("/socialmedias")
	socialMedias.Use(middleware.AuthMiddleware)
	{
		socialMedias.POST("/", smr.socialMediaController.CreateSocialMediaHandler)
		socialMedias.GET("/", smr.socialMediaController.GetSocialMediasHandler)
		socialMedias.PUT("/:uuid", smr.socialMediaController.EditSocialMediaHandler)
		socialMedias.DELETE("/:uuid", smr.socialMediaController.DeleteSocialMediaHandler)
	}
}
