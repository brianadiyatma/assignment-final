package router

import (
	"assignment-final/controller"
	"assignment-final/middleware"

	"github.com/gin-gonic/gin"
)

type photoRouter struct {
	photoController controller.PhotoController
}

func NewPhotoRouter (pc  controller.PhotoController) *photoRouter{
	return & photoRouter{
		photoController: pc,
	}
}

func (pr photoRouter) PhotoRouter (r *gin.Engine){
	photos := r.Group("/photos")
	photos.Use(middleware.AuthMiddleware)
	{
		photos.POST("/", pr.photoController.CreatePhoto)
		photos.GET("/", pr.photoController.GetPhotosHandler)
		photos.PUT("/:uuid", pr.photoController.EditPhotoHandler)
		photos.DELETE("/:uuid", pr.photoController.DeletePhotoHandler)
	}
}