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
		products.PUT("/:id", controllers.NewEditProductController().EditProduct)
		products.DELETE("/:id", controllers.NewDeleteProductController().DeleteProduct)

		// Short Polling - Devuelve los datos cada 15s
		products.GET("/short-polling", controllers.ShortPollingProducts)

		// Long Polling - Mantiene la conexión abierta hasta detectar cambios
		products.GET("/long-polling", controllers.LongPollingProducts)
	}
}
