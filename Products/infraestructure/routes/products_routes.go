package routes

import (
	"proyecto_hex/Products/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	products := r.Group("/products")
	{
		products.POST("/", controllers.NewCreateProductController().AddProduct)
		products.GET("/", controllers.NewGetAllProductsController().GetAllProducts)
		products.PUT("/", controllers.NewEditProductController().EditProduct)
		products.DELETE("/", controllers.NewDeleteProductController().DeleteProduct)

	}
}
