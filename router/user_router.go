package router

import (
	"assignment-final/controller"

	"github.com/gin-gonic/gin"
)

type userRouter struct {
	authController controller.AuthController
}

func NewUserRouter(ac controller.AuthController) *userRouter {
	return &userRouter{
		authController: ac,
	}
}

func (ar userRouter) AuthRouter(r *gin.Engine) {
	users := r.Group("/user")
	{
		users.POST("/register", ar.authController.CreateUserHandler)
	}

}
