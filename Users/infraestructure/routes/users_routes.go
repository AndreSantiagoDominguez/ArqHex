package routes

import (
	"proyecto_hex/Users/infraestructure/adapters/http/middleware"
	"proyecto_hex/Users/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		
		users.Use(middleware.CorsMiddleware())

		users.POST("/add", controllers.NewCreateUserController().AddUser)
		users.GET("/all", controllers.NewGetAllUsersController().GetAllUsers)
		users.PUT("/:id", controllers.NewEditUserController().EditUser)
		users.DELETE("/:id", controllers.NewDeleteUserController().DeleteUser)
	}
}
