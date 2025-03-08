package routes

import(
    "proyecto_hex/Purchases/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPurchasesRoutes(r *gin.Engine){
	purchases := r.Group("/purchases")
	{
		purchases.POST("/", controllers.NewCreatePurchaseController().AddPurchase)
	}
}