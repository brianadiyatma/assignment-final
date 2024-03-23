package router

import (
	"assignment-final/controller"
	"assignment-final/middleware"

	"github.com/gin-gonic/gin"
)

type userRouter struct {
	authController controller.AuthController
	userController controller.UserController
}

func NewUserRouter(ac controller.AuthController, uc controller.UserController) *userRouter {
	return &userRouter{
		authController: ac,
		userController: uc,
	}
}

func (ar userRouter) AuthRouter(r *gin.Engine) {
	users := r.Group("/user")
	{
		users.POST("/register", ar.authController.CreateUserHandler)
		users.POST("/login", ar.authController.LoginController)
		 authenticated := users.Group("/")

        authenticated.Use(middleware.AuthMiddleware)
        {

            authenticated.PUT("/:uuid", ar.userController.UpdateUser)
            authenticated.DELETE("/:uuid", ar.userController.DeleteUserHandler)

        }
		
	}
}
